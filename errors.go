package seqlease

import (
	"errors"
)

var (
	errInternalError = errors.New("Internal Error")
	errNilFunc       = errors.New("Nil Func")
	errNilReceiver   = errors.New("Nil Receiver")
)
