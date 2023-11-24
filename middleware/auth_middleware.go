package middleware

import (
	"net/http"
	"os"
	"rest_api_portfolio/helper"
	"rest_api_portfolio/model/web"
	"rest_api_portfolio/model/web/token"

	"github.com/dgrijalva/jwt-go"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}
func unauthorized(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	webResponse := web.Response{
		Code:   http.StatusUnauthorized,
		Status: "Unauthorized",
	}
	helper.WriteRequestToBody(w, webResponse)
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && (r.RequestURI == "/api/v1/auth/sign-in" || r.RequestURI == "/api/v1/auth/log-in") {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		jwtToken := r.Header.Get("Authorization")
		if jwtToken == "" {
			unauthorized(w, r)
			return
		}

		jwtTokenSecrectKey := []byte(os.Getenv("TOKEN_SECRECT_KEY"))
		claims := &token.CreateClaimsToken{}
		token, err := jwt.ParseWithClaims(jwtToken, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtTokenSecrectKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid || err == jwt.ErrHashUnavailable || err == jwt.ErrInvalidKey {
				unauthorized(w, r)
				return
			}
		}

		if !token.Valid {
			unauthorized(w, r)
			return
		}

		middleware.Handler.ServeHTTP(w, r)
	}
}
