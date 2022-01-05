package errors

type derror struct {
	Message string
	Code    uint32
}

func (e derror) Error() string {
	return e.Message
}
