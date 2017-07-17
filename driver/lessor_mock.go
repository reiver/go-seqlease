package seqleasedriver

import (
	"fmt"
)

type MockLessor struct {
	LeaseFunc func(string) (Lease, error)
}

func (receiver *MockLessor) Lease(name string) (Lease, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	fn := receiver.LeaseFunc
	if nil == fn {
		panic(fmt.Errorf("Nil Func in %T method Lease()", receiver))
	}

	return fn(name)
}
