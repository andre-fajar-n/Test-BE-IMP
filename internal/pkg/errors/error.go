// Package contains error response
package errors

import (
	"net/http"

	"github.com/go-openapi/errors"
)

func GetError(err error) errors.Error {
	if v, ok := err.(errors.Error); ok {
		return v
	}

	return errors.New(http.StatusInternalServerError, err.Error())
}

func SetError(code int32, msg string, args ...interface{}) error {
	return errors.New(code, msg, args...)
}
