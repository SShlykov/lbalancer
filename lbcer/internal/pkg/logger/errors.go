package logger

import "errors"

var (
	ErrUnknownLevel = errors.New("unknown level")
	ErrUnknownMode  = errors.New("unknown mode")
)
