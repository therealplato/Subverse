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


type ByteContent []byte

func (b *ByteContent) B() []byte {
	return []byte(b)
}

func (b *ByteContent) Type() string {
	return "application/octet-stream"
}
