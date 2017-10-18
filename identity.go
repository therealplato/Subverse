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

type NamedIdentity struct {
	Name string
}

func (i *NamedIdentity) Self() Identity {
	return i
}

func (i *NamedIdentity) Is(them Identity) bool {
	themNamed, ok := them.(*NamedIdentity)
	if !ok {
		return false
	}
	return i.Name == themNamed.Name
}
