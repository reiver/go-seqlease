package seqleasedriver

import (
	"errors"
)

var (
	errBadRequest    = errors.New("Bad Request")
	errFound         = errors.New("Found")
	errInternalError = errors.New("Internal Error")
	errNilReceiver   = errors.New("Nil Receiver")
	errNotFound      = errors.New("Not Found")
)
