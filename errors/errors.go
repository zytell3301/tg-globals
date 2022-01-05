package errors

type Derror struct {
	Message string
	Code    uint32
}

func (e Derror) Error() string {
	return e.Message
}
