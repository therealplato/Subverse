package subverse

type HTTPRef struct {
	URL string
	H   map[string]string
}

func (r *HTTPRef) URI() string {
	return r.URL
}
func (r *HTTPRef) ID() string {
	return r.URL
}
func (r *HTTPRef) Hashes() map[string]string {
	return r.H
}
