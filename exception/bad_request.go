package exception

type BadRequest struct {
	Error   error
	Message string
}

func NewBadRequest(err error, msg string) BadRequest {
	return BadRequest{
		Error:   err,
		Message: msg,
	}
}
