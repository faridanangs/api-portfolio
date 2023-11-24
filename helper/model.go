package helper

import (
	"rest_api_portfolio/model/entity"
	"rest_api_portfolio/model/web/blog"
	"rest_api_portfolio/model/web/project"
	"rest_api_portfolio/model/web/user"
)

// blog response
func BlogResponse(request entity.Blogs) blog.BlogCreateResponse {
	return blog.BlogCreateResponse{
		Id:        request.Id,
		Image:     request.Image,
		Tittle:    request.Tittle,
		Text:      request.Text,
		Name:      request.Name,
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
	}
}
func BlogsResponse(requests []entity.Blogs) []blog.BlogCreateResponse {
	var responses []blog.BlogCreateResponse

	for _, data := range requests {
		responses = append(responses, BlogResponse(data))
	}

	return responses
}

// user response
func UserResponse(request entity.Users) *user.UserCreateResponse {
	return &user.UserCreateResponse{
		IdUser:    request.IdUser,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
	}
}

// project response
func ProjectResponse(request entity.Projects) *project.CreateProjectResponse {
	return &project.CreateProjectResponse{
		Id:          request.Id,
		Tittle:      request.Tittle,
		Image:       request.Image,
		Description: request.Description,
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
	}
}
func ProjectsResponse(request []entity.Projects) []project.CreateProjectResponse {
	var projects []project.CreateProjectResponse
	for _, data := range request {
		projects = append(projects, *ProjectResponse(data))
	}
	return projects
}
