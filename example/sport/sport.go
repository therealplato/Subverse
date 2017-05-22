package sport

import "github.com/therealplato/subverse"

// Player plays the sport
type Player subverse.Identity

type Game struct {
	Name    string
	Ruleset subverse.Ref
}

var officialRef = &subverse.URIRef{
	URI: "http://www.foosball.com/learn/rules/onepage/onepage.pdf",
}
var officialRule = subverse.ReferencedRule{
	Label: "Office Foosball Rules",
	Ref:   officialRef,
	J:     "office",
}

var officeRules = []subverse.Rule{
	officialRule,
	subverse.SimpleRule{
		T: "official rules that a majority of players claim to not know shall not apply",
		J: "office",
	},
}
