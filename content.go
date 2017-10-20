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

// type RefMaker interface {
// 	MakeRef([]byte) Ref
// }

// Content is something that can become bytes
type Content interface {
	Bytes() []byte
}

// Metadata is key/values about a Ref
type Metadata interface {
	Describing() Ref
	Map() map[string]string
}

// MetaContent is something with both Metadata and Content
type MetaContent interface {
	Metadata
	Content
}

// NilContent is a placeholder
type NilContent struct{}

// Bytes returns nil
func (c NilContent) Bytes() []byte {
	return nil
}

// URIRef references a URI
type URIRef struct {
	URI   string
	bytes []byte
}

// Hash is unimplemented
func (r *URIRef) Hash() string {
	return ""
}

// Content returns the ref's bytes or nil
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

// ByteContent is a blob
type ByteContent []byte

// Bytes is the blob
func (b *ByteContent) Bytes() []byte {
	return []byte(*b)
}

// SignedContent is content plus some kind of signature
type SignedContent struct {
	Content   Content
	Signature interface{}
}
