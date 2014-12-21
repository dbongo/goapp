package errors

import "fmt"

// HTTP ...
type HTTP struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
	URL     string `json:"url"`
}

func (e *HTTP) Error() string {
	return e.Message
}

// ValidationError is an error implementation used whenever a validation failure occurs.
type ValidationError struct {
	Message string
}

func (err *ValidationError) Error() string {
	return err.Message
}

// ConflictError ...
type ConflictError ValidationError

func (err *ConflictError) Error() string {
	return err.Message
}

// NotAuthorizedError ...
type NotAuthorizedError ValidationError

func (err *NotAuthorizedError) Error() string {
	return err.Message
}

// CompositeError ...
type CompositeError struct {
	Base    error
	Message string
}

func (err *CompositeError) Error() string {
	if err.Base == nil {
		return err.Message
	}
	return fmt.Sprintf("%s Caused by: %s", err.Message, err.Base.Error())
}
