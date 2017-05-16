package subverse

type Asker interface {
}

type Query struct {
	// I'm looking for a link (photo/js/binary/csv...)
	Accept  string // "Is there a browser based tool where I can design terminal/editor color palettes?"
	Text    string
	ReplyTo Identity
}
