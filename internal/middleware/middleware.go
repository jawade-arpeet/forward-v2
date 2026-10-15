package middleware

import "forward/internal/repository"

type Middleware struct {
	Request *RequestMiddleware
}

func New(repo *repository.Repository) *Middleware {
	return &Middleware{
		Request: newRequestMiddleware(),
	}
}
