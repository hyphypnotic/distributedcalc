package transport

import (
	"context"
	"distributedcalc/backend/internal/agent/service"
	"log/slog"
)

func AgentHandler(userService service.UserService) {
	ctx := context.Background()
	pubsub := userService.Client.Subscribe(ctx, "calculate")
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			slog.Error("Error in receiving message: ", err)
			continue
		}
		expression := string(msg.Payload)
		result, err := service.EvaluateExpression(expression)
		if err != nil {
			slog.Error("Error in evaluation expression: ", err)
			continue
		}
		err = userService.Client.Publish(ctx, "calculate", result).Err()
		if err != nil {
			slog.Error("Error in publishing message to orchestrator: ", err)
		}
	}
}
