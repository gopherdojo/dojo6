package error

import "fmt"

type ApiError struct {
	message string
	code    string
}

func New(m, c string) error {
	return &ApiError{message: m, code: c}
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("Error has occurred! message = %v, code = %v", e.message, e.code)
}
