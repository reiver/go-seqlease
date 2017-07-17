package seqleasedriver

// Registrar represents a thing where (new) LessorCreates can be registered
// under a (unique) name, and from which Lessors can be created (using the
// registered LessorCreators).
type Registrar interface {
	Lessor(string, ...interface{}) (Lessor, error)
	Register(string, LessorCreator) error
}
