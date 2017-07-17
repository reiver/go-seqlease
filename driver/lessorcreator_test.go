package seqleasedriver

import (
	"testing"
)

func LessorCreatorMockLessor(t *testing.T) {

	var x LessorCreator = new(MockLessorCreator) // THIS IS WHAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}
