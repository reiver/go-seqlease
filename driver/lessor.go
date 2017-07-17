package seqleasedriver

// Lessor represents something that provides Leases.
type Lessor interface {
	Lease(string) (Lease, error)
}
