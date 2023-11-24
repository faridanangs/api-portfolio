package authcontroller

import (
	"net/http"
	"rest_api_portfolio/helper"
	"rest_api_portfolio/model/web"
	"rest_api_portfolio/model/web/token"
	"rest_api_portfolio/model/web/user"
	authservice "rest_api_portfolio/service/auth-service"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerIplm struct {
	AuthService authservice.AuthService
}

func NewAuthControllerIplm(authService authservice.AuthService) AuthController {
	return &AuthControllerIplm{
		AuthService: authService,
	}
}

func (controller *AuthControllerIplm) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var userCreateRequest user.UserCreateRequest
	helper.ReadRequestToBody(r, &userCreateRequest)

	response := controller.AuthService.Create(r.Context(), userCreateRequest)

	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}
	helper.WriteRequestToBody(w, webResponse)

}
func (controller *AuthControllerIplm) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("auth_id")

	controller.AuthService.Delete(r.Context(), userId)

	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
	}
	helper.WriteRequestToBody(w, webResponse)
}
func (controller *AuthControllerIplm) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("auth_id")

	response := controller.AuthService.FindById(r.Context(), userId)

	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}
	helper.WriteRequestToBody(w, webResponse)
}
func (controller *AuthControllerIplm) CreateToken(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	createAuthToken := token.CreateAuthToken{}
	helper.ReadRequestToBody(r, &createAuthToken)

	response := controller.AuthService.CreateToken(r.Context(), createAuthToken)

	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}
	helper.WriteRequestToBody(w, webResponse)
}
func (controller *AuthControllerIplm) RefreshToken(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	jwtToken := r.Header.Get("Authorization")

	response := controller.AuthService.RefreshToken(r.Context(), jwtToken)

	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}
	helper.WriteRequestToBody(w, webResponse)
}
