package subverse

type Identity interface {
	Self() Identity
	Is(Identity) bool
}

type Contact interface {
	ContactEndpoints() []string
}
