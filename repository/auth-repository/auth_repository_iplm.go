package authrepository

import (
	"context"
	"database/sql"
	"errors"
	"rest_api_portfolio/helper"
	"rest_api_portfolio/model/entity"
)

type AuthRepositoryIplm struct{}

func NewAuthRepositoryIplm() AuthRepository {
	return &AuthRepositoryIplm{}
}

func (repository *AuthRepositoryIplm) Save(ctx context.Context, tx *sql.Tx, request entity.Users) entity.Users {
	SQL := "INSERT INTO users(id_user, first_name, last_name, email, password, created_at, updated_at) values(?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		request.IdUser,
		request.FirstName,
		request.LastName,
		request.Email,
		request.Password,
		request.CreatedAt,
		request.UpdatedAt,
	)
	helper.PanicIfError(err, "Error ExecContext on auth-repository at Save")

	return request

}
func (repository *AuthRepositoryIplm) Delete(ctx context.Context, tx *sql.Tx, request entity.Users) {
	SQL := "DELETE FROM users WHERE id_user=?"
	_, err := tx.ExecContext(ctx, SQL, request.IdUser)
	helper.PanicIfError(err, "Error ExecContext on auth-repository at Delete")
}
func (repository *AuthRepositoryIplm) FindById(ctx context.Context, tx *sql.Tx, requestId string) (entity.Users, error) {
	SQL := "SELECT * FROM users WHERE id_user=?"
	row, err := tx.QueryContext(ctx, SQL, requestId)
	helper.PanicIfError(err, "Error QueryContext on auth-repository at FindById")
	defer row.Close()

	var user entity.Users
	if row.Next() {
		err := row.Scan(
			&user.IdUser,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		helper.PanicIfError(err, "Error row.Scan on auth-repository at FindById")
		return user, nil
	} else {
		return user, errors.New("Not Found")
	}
}
func (repository *AuthRepositoryIplm) FindByEmail(ctx context.Context, tx *sql.Tx, requestEmail string) (entity.Users, error) {
	SQL := "SELECT * FROM users WHERE email=?"
	row, err := tx.QueryContext(ctx, SQL, requestEmail)
	helper.PanicIfError(err, "Error QueryContext on auth-repository at FindByEmail")

	defer row.Close()

	var user entity.Users
	if row.Next() {
		err := row.Scan(
			&user.IdUser,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		helper.PanicIfError(err, "ErrorExecContextt onrepositoryr at CommitOrRollback")
		return user, nil
	} else {
		return user, errors.New("Not Found")
	}
}
