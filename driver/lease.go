package seqleasedriver

// Lease represents a lease oriented lock.
type Lease interface {
	Name() (string, error)
	Sequence() (int64, error)
}

