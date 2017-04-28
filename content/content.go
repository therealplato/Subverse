package content

// A Ref describes a handle that points at content
type Ref interface {
	URI() string
	ID() string
	Hashes() map[string]string
}

type Content interface {
	B() []byte
	Meta() Metadata
}

type Metadata interface {
	Content() Ref
	B() []byte
}

Type MetaContent interface{
	Metadata
	Content
}