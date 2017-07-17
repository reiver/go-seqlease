package seqleasedriver

import (
	"fmt"
)

type MockLessorCreator struct {
	LessorFunc func(...interface{}) (Lessor, error)
}

func (receiver MockLessorCreator) Lessor(params ...interface{}) (Lessor, error) {
	fn := receiver.LessorFunc

	if nil == fn {
		panic(fmt.Errorf("Nil Func in %T method Lessor()", receiver))
	}

	return fn(params...)
}
