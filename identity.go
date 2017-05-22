package subverse

type Identity interface {
	Self() Identity
	Is(Identity) bool
}

type Contact interface {
	ContactEndpoints() []string
}

// Anonymous is the identity of no one
type Anonymous struct{}

func (i *Anonymous) Self() Identity {
	return i
}
func (i *Anonymous) Is(Identity) bool {
	return false
}
