package projectrepository

import (
	"context"
	"database/sql"
	"errors"
	"rest_api_portfolio/helper"
	"rest_api_portfolio/model/entity"
)

type ProjectRepositoryIplm struct{}

func NewProjectRepositoryIplm() ProjectRepository {
	return &ProjectRepositoryIplm{}
}

func (repository *ProjectRepositoryIplm) Save(ctx context.Context, tx *sql.Tx, request entity.Projects) entity.Projects {
	SQL := "INSERT INTO projects(id, tittle, image, description, created_at, updated_at) values (?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL,
		request.Id,
		request.Tittle,
		request.Image,
		request.Description,
		request.CreatedAt,
		request.UpdatedAt,
	)
	helper.PanicIfError(err, "Error ExecContext on project-repository at Save")

	return request
}
func (repository *ProjectRepositoryIplm) Update(ctx context.Context, tx *sql.Tx, request entity.Projects) entity.Projects {
	SQL := "UPDATE projects set tittle=?, image=?, description=?, updated_at=? where id=?"
	_, err := tx.ExecContext(ctx, SQL,
		request.Tittle,
		request.Image,
		request.Description,
		request.UpdatedAt,
		request.Id,
	)
	helper.PanicIfError(err, "Error ExecContext on project-repository at Update")
	return request
}
func (repository *ProjectRepositoryIplm) Delete(ctx context.Context, tx *sql.Tx, request entity.Projects) {
	SQL := "delete from projects where id=?"
	_, err := tx.ExecContext(ctx, SQL, request.Id)
	helper.PanicIfError(err, "Error ExecContext on project-repository at Delete")
}
func (repository *ProjectRepositoryIplm) FindById(ctx context.Context, tx *sql.Tx, requestId string) (entity.Projects, error) {
	SQL := "SELECT * FROM projects where id=?"
	row, err := tx.QueryContext(ctx, SQL, requestId)
	helper.PanicIfError(err, "Error QueryContext on project-repository at FindById")
	defer row.Close()

	var project entity.Projects
	if row.Next() {
		err := row.Scan(
			&project.Id,
			&project.Tittle,
			&project.Image,
			&project.Description,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		helper.PanicIfError(err, "Error row.Scan on project-repository at FindById")
		return project, nil
	} else {
		return project, errors.New("Not Found Project")
	}
}
func (repository *ProjectRepositoryIplm) FindAll(ctx context.Context, tx *sql.Tx) []entity.Projects {
	SQL := "SELECT * FROM projects"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err, "Error QueryContext on project-repository at FindAll")
	defer rows.Close()

	var projects []entity.Projects
	for rows.Next() {
		var project entity.Projects
		err := rows.Scan(
			&project.Id,
			&project.Tittle,
			&project.Image,
			&project.Description,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		helper.PanicIfError(err, "Error row.Scan on project-repository at FindAll")
		projects = append(projects, project)
	}
	return projects
}
