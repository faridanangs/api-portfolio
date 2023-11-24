package token

type CreateAuthToken struct {
	Email    string `validate:"email,required,max=200" json:"email"`
	Password string `validate:"required" json:"password"`
}
