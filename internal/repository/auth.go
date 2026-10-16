package repository

import (
	"context"
	"forward/internal/client"
	"forward/internal/dao"
	"uuid"

	"github.com/jackc/pgx/v5"
)

type AuthRepository struct {
	pg *client.PostgresClient
}

func newAuthRepository(pg *client.PostgresClient) *AuthRepository {
	return &AuthRepository{pg: pg}
}

func (r *AuthRepository) CreateAccount(
	ctx context.Context,
	email string,
	passwordHash string,
) (*uuid.UUID, error) {
	query := `
		INSERT INTO accounts (
			email,
		 	password_hash
		)
		VALUES (
			@email,
		 	@password_hash
		)
		RETURNING id
	`
	args := pgx.NamedArgs{
		"email":         email,
		"password_hash": passwordHash,
	}

	return r.pg.QueryOne[uuid.UUID](ctx, query, args)
}

func (r *AuthRepository) GetAccountByEmail(
	ctx context.Context,
	email string,
) (*dao.AccountInfo, error) {
	query := `
		SELECT
			id,
			email,
			password_hash,
			is_active,
			is_email_verified
		FROM
			accounts
		WHERE
			email = @email
	`
	args := pgx.NamedArgs{"email": email}

	return r.pg.QueryOne[dao.AccountInfo](ctx, query, args)
}
