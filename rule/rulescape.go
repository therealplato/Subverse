package subverse

// Rule is words governing behavior
type Rule interface {
	Text() string
	Jurisdiction() string
}

// Rulescape is a collection of rules
type Rulescape interface {
	Ref() Ref
	Rules() []Rule
}

// SimpleRule is just words
type SimpleRule struct {
	T string
	J string
}

func (r SimpleRule) Text() string {
	return r.T
}
func (r SimpleRule) Jurisdiction() string {
	return r.J
}

var _ Rule = SimpleRule{}

var noBoysAllowed = SimpleRule{
	T: "No Boys Allowed",
	J: "Alice's Clubhouse",
}

// ReferencedRule points at external content
type ReferencedRule struct {
	Label string
	Ref   Ref
	J     string
}

func (r ReferencedRule) Text() string {
	c := r.Ref.Content()
	return string(c.Bytes())
}
func (r ReferencedRule) Jurisdiction() string {
	return r.J
}

var _ Rule = ReferencedRule{}
