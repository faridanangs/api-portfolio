package blogservice

import (
	"context"
	"database/sql"
	"io"
	"os"
	"rest_api_portfolio/exception"
	"rest_api_portfolio/helper"
	"rest_api_portfolio/model/entity"
	"rest_api_portfolio/model/web/blog"
	blogrepository "rest_api_portfolio/repository/blog-repository"
	"rest_api_portfolio/utils"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type BlogServiceIplm struct {
	Validate       *validator.Validate
	DB             *sql.DB
	BlogRepository blogrepository.BlogRepository
}

func NewBlogServiceIplm(validate *validator.Validate, db *sql.DB, blogRepository blogrepository.BlogRepository) BlogService {
	return &BlogServiceIplm{
		Validate:       validate,
		DB:             db,
		BlogRepository: blogRepository,
	}
}

func (service *BlogServiceIplm) Create(ctx context.Context, request blog.BlogCreateRequest) blog.BlogCreateResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err, "Error Validate.Struct on blog-service at Create")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on blog-service at Create")
	defer helper.CommitOrRollback(tx)

	data := request.FormData
	file, err := data.Open()
	helper.PanicIfError(err, "Error data.Open on blog-service at create")
	defer file.Close()

	tempFile, err := os.CreateTemp("public/blog", "image-*.png")
	helper.PanicIfError(err, "Error os.CreateTemp on blog-service at create")
	defer tempFile.Close()

	fileByte, err := io.ReadAll(file)
	helper.PanicIfError(err, "Error io.ReadAll on blog-service at create")

	tempFile.Write(fileByte)

	nameFile := tempFile.Name()

	newNameFile := strings.Split(nameFile, "\\")

	blogRequest := entity.Blogs{
		Id:        utils.Uuid(),
		Image:     newNameFile[1],
		Tittle:    request.Tittle,
		Text:      request.Text,
		Name:      request.Name,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	blogResponse := service.BlogRepository.Save(ctx, tx, blogRequest)

	return helper.BlogResponse(blogResponse)

}
func (service *BlogServiceIplm) Update(ctx context.Context, request blog.BlogUpdateRequest) blog.BlogCreateResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err, "Error Validate.Struct on blog-service at Update")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on blog-service at Update")
	defer helper.CommitOrRollback(tx)

	blogResponseByid, err := service.BlogRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFound(err, "Error FindById at Upate service blogs"))
	}

	os.Remove("public/blog/" + blogResponseByid.Image)

	data := request.FormData
	file, err := data.Open()
	helper.PanicIfError(err, "Error data.Open on blog-service at Update")
	defer file.Close()

	tempFile, err := os.CreateTemp("public/blog", "image-*.png")
	helper.PanicIfError(err, "Error os.CreateTemp on blog-service at Update")
	defer tempFile.Close()

	fileByte, err := io.ReadAll(file)
	helper.PanicIfError(err, "Error io.ReadAll on blog-service at Update")

	tempFile.Write(fileByte)
	nameFile := tempFile.Name()
	newNameFile := strings.Split(nameFile, "\\")

	blogResponseByid.Image = newNameFile[1]
	blogResponseByid.Tittle = request.Tittle
	blogResponseByid.Text = request.Text
	blogResponseByid.UpdatedAt = time.Now().Unix()

	blogResponse := service.BlogRepository.Update(ctx, tx, blogResponseByid)

	return helper.BlogResponse(blogResponse)

}
func (service *BlogServiceIplm) Delete(ctx context.Context, requestId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on project-service at Delete")
	defer helper.CommitOrRollback(tx)

	blogResponseByid, err := service.BlogRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFound(err, "Error FindById at Delete service blogs"))
	}

	service.BlogRepository.Delete(ctx, tx, blogResponseByid)
	os.Remove("public/blog/" + blogResponseByid.Image)

}
func (service *BlogServiceIplm) FindById(ctx context.Context, requestId string) blog.BlogCreateResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on project-service at FindByd")
	defer helper.CommitOrRollback(tx)

	blogResponseByid, err := service.BlogRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFound(err, "Error FindById at FindById service blogs"))
	}

	return helper.BlogResponse(blogResponseByid)
}
func (service *BlogServiceIplm) FindAll(ctx context.Context) []blog.BlogCreateResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on project-service at FindAll")
	defer helper.CommitOrRollback(tx)

	blogResponseByid := service.BlogRepository.FindAll(ctx, tx)

	return helper.BlogsResponse(blogResponseByid)
}
