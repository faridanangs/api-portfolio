package projectservice

import (
	"context"
	"rest_api_portfolio/model/web/project"
)

type ProjectService interface {
	Create(ctx context.Context, request project.CreateAndUpdateProjectRequest) project.CreateProjectResponse
	Update(ctx context.Context, request project.CreateAndUpdateProjectRequest) project.CreateProjectResponse
	Delete(ctx context.Context, requestId string)
	FindById(ctx context.Context, requestId string) project.CreateProjectResponse
	FindAll(ctx context.Context) []project.CreateProjectResponse
}
