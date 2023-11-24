package exception

type Unauthorized struct {
	Error   error
	Message string
}

func NewUnauthorized(err error, msg string) Unauthorized {
	return Unauthorized{
		Error:   err,
		Message: msg,
	}
}
