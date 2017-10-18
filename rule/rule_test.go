package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/therealplato/subverse"
)

func TestSimpleRuleDoesNotGovernAnonymous(t *testing.T) {
	soberoctober := SimpleRule{
		Ruler:       subverse.Anonymous,
		Subject:     subverse.Anonymous,
		Context:     "october 2017",
		Behavior:    "Do not consume psychoactives, except caffeine. Do not play video games.",
		Consequence: "feel shame",
	}

	assert.False(t, soberoctober.Governs(subverse.Anonymous, "october 2017"))
}

func TestSimpleRuleGovernsSubjectWithinContext(t *testing.T) {
	me := &subverse.NamedIdentity{"plato"}
	subject := me
	soberoctober := SimpleRule{
		Subject: subject,
		Context: "october 2017",
	}

	assert.True(t, soberoctober.Governs(me, "october 2017"))
}

func TestSimpleRuleDoesNotGovernSubjectOutsideContext(t *testing.T) {
	me := &subverse.NamedIdentity{"plato"}
	subject := me
	soberoctober := SimpleRule{
		Subject: subject,
		Context: "november 2017",
	}

	assert.False(t, soberoctober.Governs(me, "october 2017"))
}

func TestSimpleRuleDoesNotGovernNonSubjects(t *testing.T) {
	me := &subverse.NamedIdentity{"plato"}
	soberoctober := SimpleRule{
		Subject: &subverse.NamedIdentity{"Cat."},
		Context: "october 2017",
	}

	assert.False(t, soberoctober.Governs(me, "october 2017"))
}
