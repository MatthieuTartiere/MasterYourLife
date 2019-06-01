package httptools

import (
	log "github.com/sirupsen/logrus"
)

type HTTPError struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"message"`
	Error        string `json:"error"`
}

func NewHTTPError(msg string) HTTPError {
	return HTTPError{
		ErrorMessage: msg,
	}
}

func (this HTTPError) WithError(err error) HTTPError {
	this.Error = err.Error()
	return this
}

func (this HTTPError) WithErrorCode(code int) HTTPError {
	this.ErrorCode = code
	return this
}

func (this HTTPError) Log() HTTPError {
	log.WithFields(log.Fields{
		"error_code": this.ErrorCode,
		"error":      this.Error,
	}).Error(this.ErrorMessage)
	return this
}
