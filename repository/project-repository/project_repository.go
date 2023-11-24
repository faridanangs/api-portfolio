package projectrepository

import (
	"context"
	"database/sql"
	"rest_api_portfolio/model/entity"
)

type ProjectRepository interface {
	Save(ctx context.Context, tx *sql.Tx, request entity.Projects) entity.Projects
	Update(ctx context.Context, tx *sql.Tx, request entity.Projects) entity.Projects
	Delete(ctx context.Context, tx *sql.Tx, request entity.Projects)
	FindById(ctx context.Context, tx *sql.Tx, requestId string) (entity.Projects, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Projects
}
