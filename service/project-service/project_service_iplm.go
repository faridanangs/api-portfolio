package projectservice

import (
	"context"
	"database/sql"
	"io"
	"os"
	"rest_api_portfolio/exception"
	"rest_api_portfolio/helper"
	"rest_api_portfolio/model/entity"
	"rest_api_portfolio/model/web/project"
	projectrepository "rest_api_portfolio/repository/project-repository"
	"rest_api_portfolio/utils"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type ProjectServiceIplm struct {
	DB                *sql.DB
	Validate          *validator.Validate
	ProjectRepository projectrepository.ProjectRepository
}

func NewProjectServiceIplm(validate *validator.Validate, db *sql.DB, projectRepository projectrepository.ProjectRepository) ProjectService {
	return &ProjectServiceIplm{
		DB:                db,
		Validate:          validate,
		ProjectRepository: projectRepository,
	}
}

func (service *ProjectServiceIplm) Create(ctx context.Context, request project.CreateAndUpdateProjectRequest) project.CreateProjectResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err, "Error Validate.Struct on project-service at Create")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on project-service at Create")
	defer helper.CommitOrRollback(tx)

	data := request.FormData
	file, err := data.Open()
	helper.PanicIfError(err, "Error data.Open on project-service at Create")
	defer file.Close()

	tempFile, err := os.CreateTemp("public/project", "project-image-*.png")
	helper.PanicIfError(err, "Error os.CreateTemp on project-service at Create")
	defer tempFile.Close()

	fileByte, err := io.ReadAll(file)
	helper.PanicIfError(err, "Error io.ReadAll on project-service at Create")

	tempFile.Write(fileByte)

	fileName := tempFile.Name()

	newFileName := strings.Split(fileName, "\\")

	project := entity.Projects{
		Id:          utils.Uuid(),
		Tittle:      request.Tittle,
		Image:       newFileName[1],
		Description: request.Description,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	service.ProjectRepository.Save(ctx, tx, project)

	return *helper.ProjectResponse(project)
}
func (service *ProjectServiceIplm) Update(ctx context.Context, request project.CreateAndUpdateProjectRequest) project.CreateProjectResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err, "Error Validate.Struct on project-service at Update")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on project-service at Update")
	defer helper.CommitOrRollback(tx)

	responseById, err := service.ProjectRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFound(err, "Error FindById at Update service"))
	}
	os.Remove("public/project/" + responseById.Image)

	data := request.FormData
	file, err := data.Open()
	helper.PanicIfError(err, "Error data.Open on project-service at Update")
	defer file.Close()

	tempFile, err := os.CreateTemp("public/project", "project-image-*.png")
	helper.PanicIfError(err, "Error os.CreateTemp on project-service at Update")
	defer tempFile.Close()

	fileByte, err := io.ReadAll(file)
	helper.PanicIfError(err, "Error io.ReadAll on project-service at Update")

	tempFile.Write(fileByte)
	fileName := tempFile.Name()

	newFileName := strings.Split(fileName, "\\")

	responseById.Tittle = request.Tittle
	responseById.Image = newFileName[1]
	responseById.Description = request.Description
	responseById.UpdatedAt = time.Now().Unix()

	response := service.ProjectRepository.Update(ctx, tx, responseById)
	return *helper.ProjectResponse(response)

}
func (service *ProjectServiceIplm) Delete(ctx context.Context, requestId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on project-service at Delete")
	defer helper.CommitOrRollback(tx)

	responseById, err := service.ProjectRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFound(err, "Error FindById at Delete service"))
	}

	service.ProjectRepository.Delete(ctx, tx, responseById)
	os.Remove("public/project/" + responseById.Image)
}
func (service *ProjectServiceIplm) FindById(ctx context.Context, requestId string) project.CreateProjectResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on project-service at FindById")
	defer helper.CommitOrRollback(tx)

	responseById, err := service.ProjectRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFound(err, "Error FindById at Delete service"))
	}

	return *helper.ProjectResponse(responseById)
}
func (service *ProjectServiceIplm) FindAll(ctx context.Context) []project.CreateProjectResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on project-service at FindAll")
	defer helper.CommitOrRollback(tx)

	responses := service.ProjectRepository.FindAll(ctx, tx)

	return helper.ProjectsResponse(responses)
}
