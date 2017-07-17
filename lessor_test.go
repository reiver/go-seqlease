package seqlease

import (
	"testing"
)

func TestLessorAdaptor(t *testing.T) {
	var x Lessor = new(lessorAdaptor) // THIS IS WHAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}
