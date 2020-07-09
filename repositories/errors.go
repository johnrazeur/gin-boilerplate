package repositories

import (
	"errors"
)

// Repository errors
var (
	ErrKeyConflict   = errors.New("Key Conflict")
	ErrDataNotFound  = errors.New("Record Not Found")
	ErrUserExists    = errors.New("User already exists")
	ErrWrongPassword = errors.New("Wrong password")
)
