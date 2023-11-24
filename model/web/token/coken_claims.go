package token

import "github.com/dgrijalva/jwt-go"

type CreateClaimsToken struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	jwt.StandardClaims
}
