package blogservice

import (
	"context"
	"rest_api_portfolio/model/web/blog"
)

type BlogService interface {
	Create(ctx context.Context, request blog.BlogCreateRequest) blog.BlogCreateResponse
	Update(ctx context.Context, request blog.BlogUpdateRequest) blog.BlogCreateResponse
	Delete(ctx context.Context, requestId string)
	FindById(ctx context.Context, requestId string) blog.BlogCreateResponse
	FindAll(ctx context.Context) []blog.BlogCreateResponse
}
