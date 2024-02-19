package app

import (
	"distributedcalc/backend/internal/orchestrator/service"
	server "distributedcalc/backend/internal/orchestrator/transport"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func Run() {
	r := mux.NewRouter()
	store := sessions.NewCookieStore([]byte("dp%sco2%sa[2mni12zpmy%%vqf_w!enk"))
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	userService := service.NewUserService(store, *client)
	server.OrchestratorHandler(r, "/", userService)
	http.ListenAndServe(":8080", r)
}
