package seqleasedriver

import (
	"testing"
)

func TestMockLessor(t *testing.T) {

	var x Lessor = new(MockLessor) // THIS IS WHAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}
