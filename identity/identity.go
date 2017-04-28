package identity
type Identity interface{
	Is(Identity) bool
}

type Contact interface{
	ContactEndpoints() []string
}
