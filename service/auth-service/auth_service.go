package authservice

import (
	"context"
	"rest_api_portfolio/model/web/token"
	"rest_api_portfolio/model/web/user"
)

type AuthService interface {
	Create(ctx context.Context, request user.UserCreateRequest) user.UserCreateResponse
	Delete(ctx context.Context, requestId string)
	FindById(ctx context.Context, requestId string) user.UserCreateResponse
	CreateToken(ctx context.Context, request token.CreateAuthToken) token.TokenResponse
	RefreshToken(ctx context.Context, request string) token.TokenResponse
}
