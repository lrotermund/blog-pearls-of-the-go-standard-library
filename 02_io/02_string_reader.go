type Reader struct {
	s        string
	i        int64
	prevRune int
}

func NewReader(s string) *Reader { return &Reader{s, 0, -1} }

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}
