package seqleasedriver

import (
	"testing"
)

func TestRegistrarInternalRegistrar(t *testing.T) {

	var x Registrar = new(internalRegistrar) // THIS IS WHAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}
