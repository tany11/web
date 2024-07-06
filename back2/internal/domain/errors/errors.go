package errors

import "errors"

var (
	ErrCustomerNotFound   = errors.New("customer not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
