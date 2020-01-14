package ringbuf

type RingBuf struct {
	data  []byte
	start int
	size  int
}

func New(size int) *RingBuf {
	return &RingBuf{
		data:  make([]byte, size),
		start: 0,
		size:  size,
	}
}

func (r *RingBuf) GetSize() int {
	return r.size
}

func (r *RingBuf) Write(data []byte) {
	if r.start+len(data) < r.size {

	}
}
