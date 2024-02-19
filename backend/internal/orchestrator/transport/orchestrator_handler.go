package transport

import (
	"context"
	"distributedcalc/backend/internal/orchestrator/service"
	"distributedcalc/backend/pkg/htmlutil"
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

func OrchestratorHandler(r *mux.Router, path string, userService *service.UserService) {
	r.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		session, err := userService.CookieStore.Get(r, "session")
		if err != nil {
			slog.Error("Error in getting session: ", err)
			http.Error(w, "Internal server error!", http.StatusInternalServerError)
			return
		}
		expression := r.FormValue("expression")
		ctx := context.Background()
		err = userService.Client.Publish(ctx, "calculate", expression).Err()
		if err != nil {
			slog.Error("Failed to publish message")
			http.Error(w, "Internal server error!", http.StatusInternalServerError)
			return
		}

		if err != nil {
			slog.Error("Error in calculating: ", err)
			http.Error(w, "Calculating failed", http.StatusBadRequest)
			return
		}
		pubsub := userService.Client.Subscribe(ctx, "calculate")
		defer pubsub.Close()
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			slog.Error("Error in receiving message: ", err)
			return
		}
		result := string(msg.Payload)
		fmt.Fprint(w, result)
		err = session.Save(r, w)
		if err != nil {
			slog.Error("Session saving error: ", err)
			http.Error(w, "Failed to save session!", http.StatusInternalServerError)
			return
		}
	}).Methods(http.MethodPost)

	r.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		htmlutil.RenderTemplate(w, "main.html", nil)
	}).Methods(http.MethodGet)
}
