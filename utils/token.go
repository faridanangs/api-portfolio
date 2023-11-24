package utils

import (
	"os"
	"rest_api_portfolio/exception"
	"rest_api_portfolio/model/web/token"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(request token.CreateRequestToken, value time.Duration) string {
	jwtTokenSecrectKey := []byte(os.Getenv("TOKEN_SECRECT_KEY"))
	expiredToken := time.Now().Add(value * time.Minute)

	claims := &token.CreateClaimsToken{
		Id:        request.Id,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredToken.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtTokenSecrectKey)
	if err != nil {
		panic(exception.NewBadRequest(err, "Error SignedString at CreateToken token utils"))
	}
	return tokenString

}

func RefreshToken(jwtToken string) *token.CreateClaimsToken {
	jwtTokenSecrectKey := []byte(os.Getenv("TOKEN_SECRECT_KEY"))

	claims := &token.CreateClaimsToken{}
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtTokenSecrectKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			panic(exception.NewUnauthorized(err, "Error ParseWithClaims at refreshToken token"))
		}
	}

	if !token.Valid {
		panic(exception.NewUnauthorized(err, "Error ParseWithClaims at refreshToken !tokrn.Valid token"))

	}

	return claims
}
