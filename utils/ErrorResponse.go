package utils

type ErrorResponse struct {
	Message string
	ErrCode int
}

func (er ErrorResponse) Error() string {
	return er.Message
}