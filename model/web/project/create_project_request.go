package project

import "mime/multipart"

type CreateAndUpdateProjectRequest struct {
	FormData    *multipart.FileHeader `validate:"required"`
	Id          string                `json:"id"`
	Tittle      string                `validate:"required,max=250,min=5" json:"tittle"`
	Description string                `validate:"required,max=1000,min=20" json:"description"`
}
