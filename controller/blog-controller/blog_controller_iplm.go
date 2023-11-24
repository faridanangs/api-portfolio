package blogcontroller

import (
	"net/http"
	"rest_api_portfolio/helper"
	"rest_api_portfolio/model/web"
	"rest_api_portfolio/model/web/blog"
	blogservice "rest_api_portfolio/service/blog-service"

	"github.com/julienschmidt/httprouter"
)

type BlogControllerIplm struct {
	BlogService blogservice.BlogService
}

func NewBlogControllerIplm(blogService blogservice.BlogService) *BlogControllerIplm {
	return &BlogControllerIplm{
		BlogService: blogService,
	}
}

func (controller *BlogControllerIplm) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	helper.PanicIfError(err, "Error ParseMultipartForm on blog-controller at Create")

	file, header, err := r.FormFile("image")
	helper.PanicIfError(err, "Error r.FormFile on blog-controller at Create")
	defer file.Close()

	createBlogRequest := blog.BlogCreateRequest{}
	createBlogRequest.FormData = header
	createBlogRequest.Tittle = r.FormValue("tittle")
	createBlogRequest.Text = r.FormValue("text")
	createBlogRequest.Name = r.FormValue("name")

	response := controller.BlogService.Create(r.Context(), createBlogRequest)
	webResponse := web.Response{Code: http.StatusOK, Status: "OK", Data: response}

	helper.WriteRequestToBody(w, webResponse)

}

func (controller *BlogControllerIplm) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	helper.PanicIfError(err, "Error ParseMultipartForm on blog-controller at Update")

	file, header, err := r.FormFile("image")
	helper.PanicIfError(err, "Error r.FormFile on blog-controller at Update")
	defer file.Close()

	blogUpdateRequest := blog.BlogUpdateRequest{}
	blogId := params.ByName("blog_id")
	blogUpdateRequest.Id = blogId
	blogUpdateRequest.FormData = header
	blogUpdateRequest.Tittle = r.FormValue("tittle")
	blogUpdateRequest.Text = r.FormValue("text")

	response := controller.BlogService.Update(r.Context(), blogUpdateRequest)
	webResponse := web.Response{Code: http.StatusOK, Status: "OK", Data: response}

	helper.WriteRequestToBody(w, webResponse)
}

func (controller *BlogControllerIplm) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	blogId := params.ByName("blog_id")
	controller.BlogService.Delete(r.Context(), blogId)

	webResponse := web.Response{Code: http.StatusOK, Status: "OK"}
	helper.WriteRequestToBody(w, webResponse)
}

func (controller *BlogControllerIplm) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	blogId := params.ByName("blog_id")
	response := controller.BlogService.FindById(r.Context(), blogId)

	webResponse := web.Response{Code: http.StatusOK, Status: "OK", Data: response}
	helper.WriteRequestToBody(w, webResponse)

}

func (controller *BlogControllerIplm) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := controller.BlogService.FindAll(r.Context())

	webResponse := web.Response{Code: http.StatusOK, Status: "OK", Data: response}
	helper.WriteRequestToBody(w, webResponse)
}
