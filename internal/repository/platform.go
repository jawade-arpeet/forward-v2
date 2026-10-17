package repository

import (
	"context"
	"forward/internal/client"
	"forward/internal/dao"
	"uuid"

	"github.com/jackc/pgx/v5"
)

type PlatformRepository struct {
	pg *client.PostgresClient
}

func newPlatformRepository(pg *client.PostgresClient) *PlatformRepository {
	return &PlatformRepository{pg: pg}
}

func (r *PlatformRepository) CreatePlatform(
	ctx context.Context,
	name string,
	logoURL string,
) (*uuid.UUID, error) {
	query := `
		INSERT INTO
	  		platforms (name, logo_url)
		VALUES
	  		(@name, @logo_url)
		RETURNING id
	`
	args := pgx.NamedArgs{
		"name":     name,
		"logo_url": logoURL,
	}

	return r.pg.QueryOne[uuid.UUID](ctx, query, args)
}

func (r *PlatformRepository) GetAllPlatforms(ctx context.Context) ([]dao.PlatformInfo, error) {
	query := `
		SELECT
			id,
			name,
			logo_url
		FROM
			platforms
	`

	return r.pg.QueryAll[dao.PlatformInfo](ctx, query, nil)
}
