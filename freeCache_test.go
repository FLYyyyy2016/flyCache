package flyCache

import (
	"github.com/coocood/freecache"
	"github.com/stretchr/testify/assert"
	"testing"
)

var cache *freecache.Cache

func init() {
	cache= freecache.NewCache(100 * 1024 * 1024)
}


func TestForFreeCacheUse(t *testing.T) {
	assert.NotNil(t,cache)
	for i:=0;i<1000;i++{
		err:=cache.SetInt(int64(i),[]byte("hello"+string(i)),0)
		assert.Nil(t,err)
	}
	for i:=0;i<1000;i++{
		value,err:=cache.GetInt(int64(i))
		assert.Nil(t,err)
		assert.Equal(t,value,[]byte("hello"+string(i)))
	}
}

func BenchmarkForFreeCacheUse(b *testing.B) {
	assert.NotNil(b, cache)
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			err := cache.SetInt(int64(i), []byte("hello"+string(i)),0)
			assert.Nil(b, err)
		}
		for i := 0; i < 1000; i++ {
			value, err := cache.GetInt(int64(i))
			assert.Nil(b, err)
			assert.Equal(b, value, []byte("hello"+string(i)))
		}
	}
}