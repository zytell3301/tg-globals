package errors

type Derror struct {
	Message string
	Code    uint32
}

type InternalError struct {
	Derror
}

type EntityNotFound struct {
	Derror
}

/*
 * Status code 0 means operation successful and status code 1 indicates
 * that an internal error occurred. These two codes are reserved and no function
 * must not use these codes
 */
var (
	InternalErrorOccurred = InternalError{Derror{
		Message: "An internal error occurred",
		Code:    1,
	}}

	EntityNotFoundError = EntityNotFound{
		Derror{
			Message: "Entity not found",
		},
	}
)

func (e Derror) Error() string {
	return e.Message
}
