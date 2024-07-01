package apperror

import "errors"

var (
	ErrInputScan = errors.New("error during input scanning")
	ErrBadInput  = errors.New("bad input")
)
