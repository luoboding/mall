package errors

type Error struct {
	Code    int
	message string
}

func New(message string) *Error {
	instance := &Error{
		message: message,
	}
	return instance
}

// 实现error协议
func (e Error) Error() string {
	return e.message
}
