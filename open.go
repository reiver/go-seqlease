package seqlease

import (
	"github.com/reiver/go-seqlease/driver"

	"fmt"
)

func Open(lessorCreatorName string, params ...interface{}) (Lessor, error) {

	driverLessor, err := seqleasedriver.Registry.Lessor(lessorCreatorName, params...)
	if nil != err {
		return nil, fmt.Errorf("Problem obtaining lessor for %q, because: (%T) %v", lessorCreatorName, err, err)
	}

	// Really wish Golang had duck typing, so I wouldn't have to create or use adaptors like this!
	var lessor Lessor = &lessorAdaptor{
		DriverLessor: driverLessor,
	}

	return lessor, nil
}
