package outcomex

type Error struct {
	Code    string
	Message string
}

func NewError(code, message string) Error {
	return Error{Code: code, Message: message}
}

func ValidationError(message string) Error {
	return Error{Code: "validation", Message: message}
}

func NotFoundError(message string) Error {
	return Error{Code: "not_found", Message: message}
}

func UnexpectedError(message string) Error {
	return Error{Code: "unexpected", Message: message}
}