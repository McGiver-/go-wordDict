package dictionary

type buffer struct {
	index int
	buf   [32]byte
}

func buff(c byte) *buffer {
	b := &buffer{index: 1}
	b.prepend(c)
	return b
}

func (b *buffer) prepend(c byte) {
	b.buf[len(b.buf)-b.index] = c
	b.index++
}

func (b *buffer) string() string {
	return string(b.buf[len(b.buf)-b.index+1:])
}
