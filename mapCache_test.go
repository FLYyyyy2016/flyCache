package flyCache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var mapCache map[string][]byte

func init() {
	mapCache=make(map[string][]byte)
}

func TestMapCacheUse(t *testing.T) {
	assert.NotNil(t,mapCache)
	for i :=0;i<1000;i++{
		mapCache[string(i)]=[]byte("hello"+string(i))
	}
	for i:=0;i<1000;i++{
		if result,ok:=mapCache[string(i)];ok{
			assert.Equal(t,result,[]byte("hello"+string(i)))
		}
	}
}

func BenchmarkMapCacheUse(b *testing.B) {
	assert.NotNil(b,mapCache)
	for i := 0; i < b.N; i++ {
		for i :=0;i<1000;i++{
			mapCache[string(i)]=[]byte("hello"+string(i))
		}
		for i:=0;i<1000;i++{
			if result,ok:=mapCache[string(i)];ok{
				assert.Equal(b,result,[]byte("hello"+string(i)))
			}
		}
	}
}