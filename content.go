package subverse

// A Ref describes a handle that points at content
type Ref interface {
	URI() string
	ID() string
	Hashes() map[string]string
}

type Content interface {
	Contents() []byte
}

type Metadata interface {
	Describing() Ref
	Map() map[string]string
}

type MetaContent interface {
	Metadata
	Content
}

type ByteContent []byte

func (b *ByteContent) B() []byte {
	return []byte(*b)
}

func (b *ByteContent) Type() string {
	return "application/octet-stream"
}
