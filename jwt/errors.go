package jwt

import "errors"

var (
	ErrSegmentLen   = errors.New("a token consists of three segments connected by points")
	ErrInvalidToken = errors.New("invalid token")
)
