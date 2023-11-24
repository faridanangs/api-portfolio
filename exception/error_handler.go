package exception

import (
	"net/http"
	"rest_api_portfolio/helper"
	"rest_api_portfolio/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err any) {
	if notFound(w, r, err) {
		return
	}
	if stuctValidateError(w, r, err) {
		return
	}
	if unauthorized(w, r, err) {
		return
	}
	if badRequest(w, r, err) {
		return
	}
	internalServerError(w, r, err)
}

func stuctValidateError(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(validator.InvalidValidationError)
	if ok {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		webResponse := web.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteRequestToBody(w, webResponse)
		return true
	} else {
		return false
	}

}

func notFound(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(NotFound)
	if ok {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		webResponse := web.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception,
		}

		helper.WriteRequestToBody(w, webResponse)
		return true
	} else {
		return false
	}
}
func badRequest(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(BadRequest)
	if ok {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		webResponse := web.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception,
		}

		helper.WriteRequestToBody(w, webResponse)
		return true
	} else {
		return false
	}
}
func unauthorized(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(Unauthorized)
	if ok {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		webResponse := web.Response{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   exception,
		}

		helper.WriteRequestToBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	webResponse := web.Response{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteRequestToBody(w, webResponse)
}
