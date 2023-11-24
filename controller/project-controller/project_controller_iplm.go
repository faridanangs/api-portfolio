package projectcontroller

import (
	"net/http"
	"rest_api_portfolio/exception"
	"rest_api_portfolio/helper"
	"rest_api_portfolio/model/web"
	"rest_api_portfolio/model/web/project"
	projectservice "rest_api_portfolio/service/project-service"

	"github.com/julienschmidt/httprouter"
)

type ProjectControllerIplm struct {
	ProjectService projectservice.ProjectService
}

func NewProjectControllerIplm(projectService projectservice.ProjectService) ProjectController {
	return &ProjectControllerIplm{
		ProjectService: projectService,
	}
}

func (controller *ProjectControllerIplm) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	file, header, err := r.FormFile("image")
	if err != nil {
		panic(exception.NewBadRequest(err, "Error formFile at create controller"))
	}
	defer file.Close()

	var createAndUpdateProjectRequest project.CreateAndUpdateProjectRequest
	createAndUpdateProjectRequest.FormData = header
	createAndUpdateProjectRequest.Tittle = r.FormValue("tittle")
	createAndUpdateProjectRequest.Description = r.FormValue("description")

	response := controller.ProjectService.Create(r.Context(), createAndUpdateProjectRequest)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	}
	helper.WriteRequestToBody(w, webResponse)
}
func (controller *ProjectControllerIplm) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	projectId := params.ByName("project_id")

	file, header, err := r.FormFile("image")
	if err != nil {
		panic(exception.NewBadRequest(err, "Error formFile at update controller"))
	}
	defer file.Close()

	var createAndUpdateProjectRequest project.CreateAndUpdateProjectRequest
	createAndUpdateProjectRequest.FormData = header
	createAndUpdateProjectRequest.Id = projectId
	createAndUpdateProjectRequest.Tittle = r.FormValue("tittle")
	createAndUpdateProjectRequest.Description = r.FormValue("description")

	response := controller.ProjectService.Update(r.Context(), createAndUpdateProjectRequest)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteRequestToBody(w, webResponse)
}
func (controller *ProjectControllerIplm) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	projectId := params.ByName("project_id")
	controller.ProjectService.Delete(r.Context(), projectId)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
	}
	helper.WriteRequestToBody(w, webResponse)

}
func (controller *ProjectControllerIplm) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	projectId := params.ByName("project_id")
	response := controller.ProjectService.FindById(r.Context(), projectId)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	}
	helper.WriteRequestToBody(w, webResponse)
}
func (controller *ProjectControllerIplm) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := controller.ProjectService.FindAll(r.Context())

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	}
	helper.WriteRequestToBody(w, webResponse)
}
