package seqleasedriver

import (
	"errors"
)

var (
	errFound         = errors.New("Found")
	errInternalError = errors.New("Internal Error")
	errNilReceiver   = errors.New("Nil Receiver")
	errNotFound      = errors.New("Not Found")
)
