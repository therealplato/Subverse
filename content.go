package subverse

import (
	"io/ioutil"
	"net/http"
)

// A Ref describes a handle that points at content
type Ref interface {
	Hash() string
	Content() Content
}

type RefMaker interface {
	MakeRef([]byte) Ref
}

type Content interface {
	Bytes() []byte
}

type Metadata interface {
	Describing() Ref
	Map() map[string]string
}

type MetaContent interface {
	Metadata
	Content
}

type NilContent struct{}

func (c NilContent) Bytes() []byte {
	return nil
}

type URIRef struct {
	URI   string
	bytes []byte
}

func (r *URIRef) Hash() string {
	return ""
}
func (r *URIRef) Content() Content {
	if r.bytes == nil {
		return NilContent{}
	}
	bc := ByteContent(r.bytes)
	return &bc
}
func (r *URIRef) populateContent() {
	if r.bytes == nil {
		resp, err := http.Get(r.URI)
		if err != nil {
			return
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		r.bytes = b
	}
}

type ByteContent []byte

func (b *ByteContent) Bytes() []byte {
	return []byte(*b)
}
