package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootError  error  `json:"root_error"`
	Message    string `json:"message"`
}

func NewErrorResponse(root error, msg string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootError:  root,
		Message:    msg,
	}
}

func (a *AppError) Error() string {
	return a.GetRootErr().Error()
}

func (a *AppError) GetRootErr() error {
	if err, ok := a.RootError.(*AppError); ok {
		return err.GetRootErr()
	}
	return a.RootError
}

func NewCustomError(root error, msg string) *AppError {
	if root != nil {
		return NewErrorResponse(root, root.Error())
	}
	return NewErrorResponse(errors.New(msg), msg)
}
