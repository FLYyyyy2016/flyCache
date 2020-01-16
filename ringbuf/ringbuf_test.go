package ringbuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var rb *RingBuf

func init() {
	rb = New(20)
}

func TestRingBuf_WriteRead(t *testing.T) {
	// given
	data := make([]byte, 5)
	dataOver := make([]byte, 6)
	rb = New(5)

	// then
	rb.Write([]byte("123"))
	rb.ReadAt(data, 0)

	// when
	assert.Equal(t, []byte("123"), data[:3])
	assert.Equal(t, 3, rb.writeIndex)

	// then
	rb.Write([]byte("12345"))
	rb.ReadAt(data, 0)
	rb.ReadAt(dataOver, 0)

	// when
	assert.Equal(t, []byte("34512"), data)
	assert.Equal(t, []byte("34512"), dataOver[:rb.GetSize()])
	assert.Equal(t, []byte("12345"), rb.ReadAll())
	assert.Equal(t, 3, rb.writeIndex)

	// then
	rb.WriteAt([]byte("54321"), 2)
	rb.ReadAt(data, 0)

	// when
	assert.Equal(t, []byte("43215"), rb.ReadAll())
	assert.Equal(t, []byte("21543"), data)
	assert.Equal(t, 3, rb.writeIndex)

	// then
	rb.Write([]byte("123456789"))
	rb.ReadAt(data, 0)

	// when
	assert.Equal(t, 2, rb.writeIndex)
	assert.Equal(t, []byte("56789"), rb.ReadAll())
	assert.Equal(t, []byte("89567"), data)

	// then
	rb.WriteAt([]byte("123456789"), 4)
	rb.ReadAt(data, 0)

	// when
	assert.Equal(t, 2, rb.writeIndex)
	assert.Equal(t, []byte("95678"), rb.ReadAll())
	assert.Equal(t, []byte("78956"), data)

	// then
	rb.ReadAt(data, 2)

	// when
	assert.Equal(t, data, rb.ReadAll())

}
