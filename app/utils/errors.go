package utils

import "errors"

var (
	ErrEmptySearchRequest           = errors.New("search request is empty")
	ErrNoBodyData                 = errors.New("no body data")
	ErrAuthorization              = errors.New("authorization error")
	ErrAuthorizationNotValidToken = errors.New("authorization error not valid token")
)
