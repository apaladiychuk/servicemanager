package model

import (
	"fmt"
	"net/http"
)

type (
	Errors struct {
		Code     ErrorWrap
		Message  string
		Original error
	}

	ErrorWrap int

	ErrorResponse struct {
		Status int         `json:"status,omitempty"`
		Error  interface{} `json:"error,omitempty"`
	}
)

type RestResponse struct {
	Status int         `json:"status,omitempty"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

const (
	ErrorPermission ErrorWrap = http.StatusForbidden
	NotFound        ErrorWrap = http.StatusNotFound
	Internal        ErrorWrap = http.StatusInternalServerError
	InvalidInput    ErrorWrap = http.StatusBadRequest
)

func (e Errors) Error() string {
	return e.Message
}

func (e ErrorWrap) Wrap(eo error, msg string) Errors {
	return Errors{
		Code:     e,
		Message:  msg,
		Original: eo,
	}
}

func (e ErrorWrap) WrapF(eo error, msg string, args ...interface{}) Errors {
	return Errors{
		Code:     e,
		Message:  fmt.Sprintf(msg, args...),
		Original: eo,
	}
}

func (e ErrorWrap) New(msg string) Errors {
	return Errors{
		Code:     e,
		Message:  msg,
		Original: nil,
	}
}

func (e ErrorWrap) NewF(msg string, args ...interface{}) Errors {
	return Errors{
		Code:     e,
		Message:  fmt.Sprintf(msg, args...),
		Original: nil,
	}
}
