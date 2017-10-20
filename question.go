package subverse

// Asker is unimplemented
type Asker interface {
}

// Query describes an information request
type Query struct {
	Accept  string // I'm looking for a link (photo/js/binary/csv...)
	Text    string // "Is there a browser based tool where I can design terminal/editor color palettes?"
	ReplyTo Identity
}
