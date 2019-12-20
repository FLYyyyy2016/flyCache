package flyCache

import (
	"github.com/allegro/bigcache"
	"github.com/stretchr/testify/assert"
	"testing"
)

var bigCache *bigcache.BigCache

func init() {
	bigCache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(1))
}

func TestForBigCacheUse(t *testing.T) {
	assert.NotNil(t, bigCache)
	for i := 0; i < 1000; i++ {
		err := bigCache.Set(string(i), []byte("hello"+string(i)))
		assert.Nil(t, err)
	}
	for i := 0; i < 1000; i++ {
		value, err := bigCache.Get(string(i))
		assert.Nil(t, err)
		assert.Equal(t, value, []byte("hello"+string(i)))
	}
}
func BenchmarkForBigCacheUse(b *testing.B) {
	assert.NotNil(b, bigCache)
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			err := bigCache.Set(string(i), []byte("hello"+string(i)))
			assert.Nil(b, err)
		}
		for i := 0; i < 1000; i++ {
			value, err := bigCache.Get(string(i))
			assert.Nil(b, err)
			assert.Equal(b, value, []byte("hello"+string(i)))
		}
	}
}
