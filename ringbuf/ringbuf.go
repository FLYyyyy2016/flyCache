package ringbuf

import (
	"fmt"
	"sync"
)

type RingBuf struct {
	data       []byte
	writeIndex int
	aroundTime int
	size       int
	sync.RWMutex
}

func (r *RingBuf) String() string {
	return fmt.Sprintf("data:%s\nindex is %d,size is %d\ndata len=%d,data cap = %d", string(r.data), r.writeIndex, r.size, len(r.data), cap(r.data))
}

func New(size int) *RingBuf {
	return &RingBuf{
		data: make([]byte, size),
		size: size,
	}
}

func (r *RingBuf) GetSize() int {
	return r.size
}

func (r *RingBuf) GetAroundTime() int {
	return r.aroundTime
}

func (r *RingBuf) Write(data []byte) {
	index := r.writeIndex % r.size
	if len(data) > r.size {
		index := (r.writeIndex + len(data)) % r.size
		copy(r.data[index:], data[len(data)-r.size:])
		copy(r.data[:index], data[len(data)-index:])
	} else {
		if index+len(data) < r.size {
			copy(r.data[index:], data)
		} else {
			length := copy(r.data[index:], data)
			copy(r.data[:], data[length:])
		}
	}
	r.writeIndex += len(data)

}

func (r *RingBuf) WriteAt(data []byte, index int) {
	if index > r.size {
		index = index % r.size
	}
	if len(data) > r.size {
		index = (index + len(data)) % r.size
		copy(r.data[index:], data[len(data)-r.size:])
		copy(r.data[:index], data[len(data)-index:])
	} else {
		if index+len(data) < r.size {
			copy(r.data[index:], data)
		} else {
			length := copy(r.data[index:], data)
			copy(r.data[:], data[length:])
		}
	}

}

func (r *RingBuf) ReadAt(data []byte, index int) {
	if index > r.size {
		index = index % r.size
	}
	l := len(data)
	if l > r.size {
		l = r.size
	}
	if index+l < r.size {
		copy(data, r.data[index:])
	} else {
		length := copy(data, r.data[index:])
		copy(data[length:], r.data[:index])
	}
}

func (r *RingBuf) ReadAll() []byte {
	index := r.writeIndex % r.size
	return append(r.data[index:], r.data[:index]...)
}

func (r *RingBuf) WriteIndex() int {
	return r.writeIndex
}
