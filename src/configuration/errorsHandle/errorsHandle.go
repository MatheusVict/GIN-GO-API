package errorsHandle

import "net/http"

type ErrorsHandle struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ErrorsHandle) Error() string {
	return e.Message
}

func NewErrorsHandle(message string, err string, code int, causes []Causes) *ErrorsHandle {
	return &ErrorsHandle{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *ErrorsHandle {
	return &ErrorsHandle{
		Message: message,
		Err:     "Bad request",
		Code:    http.StatusBadRequest,
	}
}
func NewBadRequestValidationError(message string, causes []Causes) *ErrorsHandle {
	return &ErrorsHandle{
		Message: message,
		Err:     "Bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewNotFoundError(message string) *ErrorsHandle {
	return &ErrorsHandle{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}

func NewInternalServerError(message string) *ErrorsHandle {
	return &ErrorsHandle{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}
