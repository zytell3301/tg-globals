package errors

type Derror struct {
	Message string
	Code    uint32
}

type InternalError struct {
	Derror
}

var (
	InternalErrorOccurred = InternalError{Derror{
		Message: "An internal error occurred",
		Code:    1,
	}}
)

func (e Derror) Error() string {
	return e.Message
}
