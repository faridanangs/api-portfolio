package authservice

import (
	"context"
	"database/sql"
	"os"
	"rest_api_portfolio/exception"
	"rest_api_portfolio/helper"
	"rest_api_portfolio/model/entity"
	"rest_api_portfolio/model/web/token"
	"rest_api_portfolio/model/web/user"
	authrepository "rest_api_portfolio/repository/auth-repository"
	"rest_api_portfolio/utils"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceIplm struct {
	Validate       *validator.Validate
	DB             *sql.DB
	AuthRepository authrepository.AuthRepository
}

func NewAuthServiceIplm(validate *validator.Validate, db *sql.DB, authRepository authrepository.AuthRepository) AuthService {
	return &AuthServiceIplm{
		Validate:       validate,
		DB:             db,
		AuthRepository: authRepository,
	}
}

func (service *AuthServiceIplm) Create(ctx context.Context, request user.UserCreateRequest) user.UserCreateResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err, "Error Validate.Struct on auth-service at Create")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on auth-service at Create")
	defer helper.CommitOrRollback(tx)

	// generate hash password
	hashPassword, err := utils.HashPassword(request.Password)

	user := entity.Users{
		IdUser:    utils.Uuid(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  hashPassword,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	response := service.AuthRepository.Save(ctx, tx, user)

	return *helper.UserResponse(response)

}
func (service *AuthServiceIplm) Delete(ctx context.Context, requestId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on auth-service at Delete")
	defer helper.CommitOrRollback(tx)

	userResponse, err := service.AuthRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFound(err, "Error FindById at Delete auth-service"))
	}
	service.AuthRepository.Delete(ctx, tx, userResponse)
}
func (service *AuthServiceIplm) FindById(ctx context.Context, requestId string) user.UserCreateResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on auth-service at FindById")
	defer helper.CommitOrRollback(tx)

	responseById, err := service.AuthRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFound(err, "Error FindById at FindById auth-service"))
	}

	return *helper.UserResponse(responseById)
}

func (service *AuthServiceIplm) CreateToken(ctx context.Context, request token.CreateAuthToken) token.TokenResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on auth-service at CreateToken")
	defer helper.CommitOrRollback(tx)

	responseByEmail, err := service.AuthRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewNotFound(err, "Error FindByEmail at CreateToken auth-service"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(responseByEmail.Password), []byte(request.Password))
	if err != nil {
		panic(exception.NewBadRequest(err, "Error CompareHashAndPassword at CreateToken auth-service"))
	}

	tokenExpired, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRED_TIME"))
	tokenExpiredRefresh, _ := strconv.Atoi(os.Getenv("TOKEN_REFRESH_EXPIRED_TIME"))

	createRequestToken := token.CreateRequestToken{
		Id:        responseByEmail.IdUser,
		FirstName: responseByEmail.FirstName,
		LastName:  responseByEmail.LastName,
		Email:     responseByEmail.Email,
	}

	tokenResponse := token.TokenResponse{
		Token:        utils.CreateToken(createRequestToken, time.Duration(tokenExpired)),
		RefreshToken: utils.CreateToken(createRequestToken, time.Duration(tokenExpiredRefresh)),
	}

	return tokenResponse
}
func (service *AuthServiceIplm) RefreshToken(ctx context.Context, request string) token.TokenResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "Error DB.Begin on auth-service at RefreshToken")
	defer helper.CommitOrRollback(tx)

	tokenExpired, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRED_TIME"))
	tokenExpiredRefresh, _ := strconv.Atoi(os.Getenv("TOKEN_REFRESH_EXPIRED_TIME"))

	claims := utils.RefreshToken(request)
	_, err = service.AuthRepository.FindById(ctx, tx, claims.Id)
	if err != nil {
		panic(exception.NewNotFound(err, "Error fincById at RefreshToken service"))
	}

	createRequestToken := token.CreateRequestToken{
		Id:        claims.Id,
		FirstName: claims.FirstName,
		LastName:  claims.LastName,
		Email:     claims.Email,
	}

	tokenResponse := token.TokenResponse{
		Token:        utils.CreateToken(createRequestToken, time.Duration(tokenExpired)),
		RefreshToken: utils.CreateToken(createRequestToken, time.Duration(tokenExpiredRefresh)),
	}

	return tokenResponse
}
