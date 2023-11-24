package authcontroller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	CreateToken(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	RefreshToken(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
