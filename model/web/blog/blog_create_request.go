package blog

import "mime/multipart"

type BlogCreateRequest struct {
	FormData *multipart.FileHeader `validate:"required"`
	Tittle   string                `validate:"required,min=10,max=300" json:"tittle"`
	Text     string                `validate:"required" json:"text"`
	Name     string                `validate:"required,min=3,max=50" json:"name"`
}
