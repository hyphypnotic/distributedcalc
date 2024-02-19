package service

import (
	"github.com/gorilla/sessions"
	"github.com/redis/go-redis/v9"
)

type UserService struct {
	CookieStore *sessions.CookieStore
	Client      redis.Client
}

func NewUserService(store *sessions.CookieStore, client redis.Client) *UserService {
	return &UserService{
		CookieStore: store,
		Client:      client,
	}
}
