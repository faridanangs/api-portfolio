package exception

type NotFound struct {
	Error   error
	Message string
}

func NewNotFound(err error, msg string) NotFound {
	return NotFound{
		Error:   err,
		Message: msg,
	}
}
