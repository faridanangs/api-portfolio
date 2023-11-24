package blogrepository

import (
	"context"
	"database/sql"
	"rest_api_portfolio/model/entity"
)

type BlogRepository interface {
	Save(ctx context.Context, tx *sql.Tx, request entity.Blogs) entity.Blogs
	Update(ctx context.Context, tx *sql.Tx, request entity.Blogs) entity.Blogs
	Delete(ctx context.Context, tx *sql.Tx, request entity.Blogs)
	FindById(ctx context.Context, tx *sql.Tx, requestId string) (entity.Blogs, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Blogs
}
