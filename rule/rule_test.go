package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/therealplato/subverse"
)

func TestSimpleRuleDoesNotGovernAnonymous(t *testing.T) {
	ruler := subverse.Anonymous
	subject := subverse.Anonymous
	soberoctober := SimpleRule{
		Ruler:    ruler,
		Subject:  subject,
		Context:  "october 2017",
		Behavior: "Do not consume psychoactives, except caffeine. Do not play video games.",
	}

	me := subverse.Anonymous
	assert.False(t, soberoctober.Governs(me, "october 2017"))
}
