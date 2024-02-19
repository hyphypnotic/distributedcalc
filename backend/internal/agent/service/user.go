package service

import (
	"github.com/gorilla/sessions"
	"github.com/redis/go-redis/v9"
)

type UserService struct {
	Client redis.Client
}

func NewUserService(store *sessions.CookieStore, client redis.Client) *UserService {
	return &UserService{
		Client: client,
	}
}
