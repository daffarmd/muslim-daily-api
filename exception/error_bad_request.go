package exception

type BadRequestError struct {
	Message string
	Errors  map[string]string
}

func NewBadRequestError(message string, errors map[string]string) BadRequestError {
	return BadRequestError{
		Message: message,
		Errors:  errors,
	}
}
