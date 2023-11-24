package blogrepository

import (
	"context"
	"database/sql"
	"errors"
	"rest_api_portfolio/helper"
	"rest_api_portfolio/model/entity"
)

type BlogRepositoryIplm struct{}

func NewBlogRepositoryIplm() BlogRepository {
	return &BlogRepositoryIplm{}
}

func (repository *BlogRepositoryIplm) Save(ctx context.Context, tx *sql.Tx, request entity.Blogs) entity.Blogs {
	SQL := "insert into blogs(id, image, tittle, text, name, created_at, updated_at) values(?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL,
		request.Id,
		request.Image,
		request.Tittle,
		request.Text,
		request.Name,
		request.CreatedAt,
		request.UpdatedAt,
	)
	helper.PanicIfError(err, "Error ExecContext on blog-repository at Save")

	return request
}
func (repository *BlogRepositoryIplm) Update(ctx context.Context, tx *sql.Tx, request entity.Blogs) entity.Blogs {
	SQL := "update blogs set image=?, tittle=?, text=?, updated_at=? where id =?"
	_, err := tx.ExecContext(ctx, SQL,
		request.Image,
		request.Tittle,
		request.Text,
		request.UpdatedAt,
		request.Id,
	)
	helper.PanicIfError(err, "Error ExecContext on blog-repository at Update")

	return request
}
func (repository *BlogRepositoryIplm) Delete(ctx context.Context, tx *sql.Tx, request entity.Blogs) {
	SQL := "delete from blogs where id =?"
	_, err := tx.ExecContext(ctx, SQL, request.Id)
	helper.PanicIfError(err, "Error ExecContext on blog-repository at Delete")

}
func (repository *BlogRepositoryIplm) FindById(ctx context.Context, tx *sql.Tx, requestId string) (entity.Blogs, error) {
	SQL := "select * from blogs where id=?"
	row, err := tx.QueryContext(ctx, SQL, requestId)
	helper.PanicIfError(err, "Error QueryContext on blog-repository at FindById")
	defer row.Close()

	var blogResponse entity.Blogs

	if row.Next() {
		err := row.Scan(
			&blogResponse.Id,
			&blogResponse.Image,
			&blogResponse.Tittle,
			&blogResponse.Text,
			&blogResponse.Name,
			&blogResponse.CreatedAt,
			&blogResponse.UpdatedAt,
		)
		helper.PanicIfError(err, "Error row.Scan on blog-repository at FindById")
		return blogResponse, nil
	} else {
		return blogResponse, errors.New("Blog Not Found")
	}
}
func (repository *BlogRepositoryIplm) FindAll(ctx context.Context, tx *sql.Tx) []entity.Blogs {
	SQL := "select * from blogs"
	row, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err, "Error QueryContext on blog-repository at FindAll")
	defer row.Close()

	var blogResponses []entity.Blogs

	for row.Next() {
		var blogResponse entity.Blogs
		err := row.Scan(
			&blogResponse.Id,
			&blogResponse.Image,
			&blogResponse.Tittle,
			&blogResponse.Text,
			&blogResponse.Name,
			&blogResponse.CreatedAt,
			&blogResponse.UpdatedAt,
		)
		helper.PanicIfError(err, "Error row.Scan on blog-repository at FindAll")
		blogResponses = append(blogResponses, blogResponse)
	}

	return blogResponses
}
