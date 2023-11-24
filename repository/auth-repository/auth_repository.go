package authrepository

import (
	"context"
	"database/sql"
	"rest_api_portfolio/model/entity"
)

type AuthRepository interface {
	Save(ctx context.Context, tx *sql.Tx, request entity.Users) entity.Users
	Delete(ctx context.Context, tx *sql.Tx, reques entity.Users)
	FindById(ctx context.Context, tx *sql.Tx, requestId string) (entity.Users, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, requestEmail string) (entity.Users, error)
}
