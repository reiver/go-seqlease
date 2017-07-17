package seqlease

import (
	"github.com/reiver/go-seqlease/driver"
)

// Really wish Golang had duck typing, so I wouldn't have to create or use adaptors like this!
type lessorAdaptor struct {
	DriverLessor seqleasedriver.Lessor
}

func (receiver *lessorAdaptor) Lease(name string) (Lease, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	driverLessor := receiver.DriverLessor
	if nil == driverLessor {
		return nil, errInternalError
	}

	return driverLessor.Lease(name)
}
