package store

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrNoRowsAffected = errors.New("no rows affected")
)
