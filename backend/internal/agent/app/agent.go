package app

import (
	"distributedcalc/backend/internal/agent/service"
	server "distributedcalc/backend/internal/agent/transport"
	"github.com/redis/go-redis/v9"
)

func Run() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	userService := service.UserService{Client: *client}
	server.AgentHandler(userService)
}
