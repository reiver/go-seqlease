package seqleasedriver

// LessorCreator represents something that creates Lessors.
type LessorCreator interface {
	Lessor(...interface{}) (Lessor, error)
}
