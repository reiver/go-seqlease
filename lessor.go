package seqlease

// Lessor represents something that provides leases.
type Lessor interface {
	Lease(string) (Lease, error)
}
