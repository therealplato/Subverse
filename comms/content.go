package content

type ByteContent []byte

func (b *ByteContent) B() []byte {
	return []byte(b)
}

func (b *ByteContent) Type() string {
	return "application/octet-stream"
}
