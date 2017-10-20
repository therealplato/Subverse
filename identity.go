package subverse

// Identity has a Self and may be other Identities
type Identity interface {
	Self() Identity
	Is(Identity) bool
}

// type Contact interface {
// 	ContactEndpoints() []string
// }

// Anonymous is no one
var Anonymous = &AnonymousIdentity{}

// AnonymousIdentity is the identity of no one
type AnonymousIdentity struct{}

// Self returns Anonymous
func (i *AnonymousIdentity) Self() Identity {
	return Anonymous
}

// Is returns false. Anonymous is no one
func (i *AnonymousIdentity) Is(Identity) bool {
	return false
}

// NamedIdentity is defined by a name
type NamedIdentity struct {
	Name string
}

// Self returns the NamedIdentity
func (i *NamedIdentity) Self() Identity {
	return i
}

// Is returns true if the names are the same
// Is returns false if the names mismatch or them is not a NamedIdentity
func (i *NamedIdentity) Is(them Identity) bool {
	themNamed, ok := them.(*NamedIdentity)
	if !ok {
		return false
	}
	return i.Name == themNamed.Name
}
