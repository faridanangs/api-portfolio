package blog

import "mime/multipart"

type BlogUpdateRequest struct {
	Id       string                `json:"id"`
	FormData *multipart.FileHeader `validate:"required"`
	Tittle   string                `validate:"required,min=10,max=300" json:"tittle"`
	Text     string                `validate:"required" json:"text"`
}
