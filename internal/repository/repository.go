package repository

import (
	"forward/internal/client"
)

type Repository struct {
	Auth     *AuthRepository
	Platform *PlatformRepository
}

func New(clt *client.Client) *Repository {
	return &Repository{
		Auth:     newAuthRepository(clt.Postgres),
		Platform: newPlatformRepository(clt.Postgres),
	}
}
