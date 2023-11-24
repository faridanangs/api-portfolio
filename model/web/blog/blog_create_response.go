package blog

type BlogCreateResponse struct {
	Id        string `json:"id"`
	Image     string `json:"image"`
	Tittle    string `json:"tittle"`
	Text      string `json:"text"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
