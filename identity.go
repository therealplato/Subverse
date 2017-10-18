package subverse

type Identity interface {
	Self() Identity
	Is(Identity) bool
}

type Contact interface {
	ContactEndpoints() []string
}

var Anonymous = &AnonymousIdentity{}

// AnonymousIdentity is the identity of no one
type AnonymousIdentity struct{}

func (i *AnonymousIdentity) Self() Identity {
	return i
}
func (i *AnonymousIdentity) Is(Identity) bool {
	return false
}
