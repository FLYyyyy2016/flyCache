package flyCache

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache_GetSet(t *testing.T) {
	// given
	cache := NewCache(1024 * 1024 * 4)
	count := 1024

	// when
	for i := 0; i < count; i++ {
		cache.Set([]byte(strconv.Itoa(i)), []byte("k:"+strconv.Itoa(i)))
	}

	// then
	for i := 0; i < count; i++ {
		assert.Equal(t, []byte("k:"+strconv.Itoa(i)), cache.Get([]byte(strconv.Itoa(i))))
	}
}
