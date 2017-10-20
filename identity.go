package subverse

import (
	"crypto"
	"crypto/ecdsa"
	"errors"
	"math/big"
)

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

var _ Identity = &AnonymousIdentity{}

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

var _ Identity = &NamedIdentity{}

// ErrMissingSignature indicates a signature should have been present and was not
var ErrMissingSignature = errors.New("The signed content had a missing signature")

// ErrInvalidSignature indicates a signature was present but was not produced by this Identity
var ErrInvalidSignature = errors.New("The signed content had an incorrect signature")

// ErrMissingHashFunc indicates the Identity used a hash function that I don't have
var ErrMissingHashFunc = errors.New("HashFunc was not available")

// CryptoIdentity is an asymetric keypair representing an Identity
type CryptoIdentity interface {
	Identity
	Verify(SignedContent) error
}

// ECDSAIdentity is an ECDSA keypair that represents an Identity
type ECDSAIdentity struct {
	Keypair  ecdsa.PrivateKey
	HashFunc crypto.Hash
}

// Self returns the ECDSAIdentity
func (i *ECDSAIdentity) Self() Identity {
	return i
}

// Is returns false if them is not a *ECDSAIdentity
// Is returns true if i and them have matching public keys
func (i *ECDSAIdentity) Is(them Identity) bool {
	j, ok := them.(*ECDSAIdentity)
	if !ok {
		return false
	}
	return i.Keypair.PublicKey == j.Keypair.PublicKey
}

// Verify returns nil if sc validates as signed by i
// Returns ErrMissingSignature if sig is missing
// Returns ErrInvalidSignature if sig is present but wrong
// Returns ErrMissingHashFunc if hash func is unavailable
func (i *ECDSAIdentity) Verify(sc SignedContent) error {
	rs, ok := sc.Signature.([]*big.Int)
	if !ok {
		return ErrMissingSignature
	}
	if len(rs) != 2 {
		return ErrMissingSignature
	}
	if rs[0] == nil || rs[1] == nil {
		return ErrMissingSignature
	}
	if !i.HashFunc.Available() {
		return ErrMissingHashFunc
	}
	hasher := i.HashFunc.New()
	hashed := hasher.Sum(sc.Content.Bytes())
	valid := ecdsa.Verify(&(i.Keypair.PublicKey), hashed, rs[0], rs[1])
	if valid != true {
		return ErrInvalidSignature
	}
	return nil
}

var _ Identity = &ECDSAIdentity{}
var _ CryptoIdentity = &ECDSAIdentity{}
