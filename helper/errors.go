package helper

import "errors"

// service errors
var (
	ErrInvalidOption             = errors.New("invalid option. expects integer [1-5]")
	ErrValidation                = errors.New("invalid request. expects integer [1-5]")
	ErrRandomUrlServiceNotPassed = errors.New("random url service url was not passed")
	ErrRandomUrlServiceInvalid   = errors.New("random url service url is invalid")
)
