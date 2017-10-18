package rule

import "github.com/therealplato/subverse"

// Rule is words governing behavior
type Rule interface {
	Governs(subject subverse.Identity, context string) bool
}

// SimpleRule governs the behavior of one subject
type SimpleRule struct {
	Ruler       subverse.Identity
	Subject     subverse.Identity
	Behavior    string
	Context     string
	Consequence string
}

// Governs decides if the rule applies to an identity in a context
func (s *SimpleRule) Governs(potentialSubject subverse.Identity, context string) bool {
	return (potentialSubject.Is(s.Subject) && context == s.Context)
}
