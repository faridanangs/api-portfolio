package user

type UserCreateRequest struct {
	FirstName string `validate:"required,max=10,min=3" json:"first_name"`
	LastName  string `validate:"max=30" json:"last_name"`
	Email     string `validate:"required,max=200,email" json:"email"`
	Password  string `validate:"required" json:"password"`
}
