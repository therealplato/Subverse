package sport

import (
	"github.com/therealplato/subverse"
	"github.com/therealplato/subverse/rule"
)

// Player plays the sport
type Player subverse.Identity

type Game struct {
	Name    string
	Ruleset subverse.Ref
}

var officialRef = &subverse.URIRef{
	URI: "http://www.foosball.com/learn/rules/onepage/onepage.pdf",
}
var officialRule = rule.ReferencedRule{
	Label: "Office Foosball Rules",
	Ref:   officialRef,
	J:     "office",
}

var officeRules = []rule.Rule{
	officialRule,
	// subverse.SimpleRule{
	// 	T: "official rules that a majority of players claim to not know shall not apply",
	// 	J: "office",
	// },
}
