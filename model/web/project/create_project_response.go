package project

type CreateProjectResponse struct {
	Id          string `json:"id"`
	Tittle      string `json:"tittle"`
	Image       string `json:"image"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
