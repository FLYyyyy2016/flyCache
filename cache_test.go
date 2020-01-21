package flyCache

import (
	"fmt"
	"math/rand"
	_ "net/http/pprof"
	"reflect"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/pingcap/go-ycsb/pkg/generator"

	"github.com/stretchr/testify/assert"
)

const workloadSize = 2 << 20

func TestCache_GetSet(t *testing.T) {
	// given
	cache := NewCache(workloadSize)
	count := workloadSize / 4

	// when
	t.Run("set count key-val", func(t *testing.T) {
		for i := 0; i < count; i++ {
			if err := cache.Set([]byte(fmt.Sprintf("key:%6d", i)), []byte(fmt.Sprintf("val:%6d", i))); err != nil {
				t.Error(err)
			}
		}
	})

	// then
	t.Run("get count key-val", func(t *testing.T) {
		errTime := 0
		for i := 0; i < count; i++ {
			data, err := cache.Get([]byte(fmt.Sprintf("key:%6d", i)))
			if err != nil {
				errTime++
			} else {
				assert.Equal(t, []byte(fmt.Sprintf("val:%6d", i)), data)
			}
		}

		assert.Equal(t, count, cache.GetCount())
		assert.Equal(t, errTime, cache.GetMiss())
	})

}

func BenchmarkGetSet(b *testing.B) {
	cache := NewCache(workloadSize)
	zipList := zipfKeyList()
	b.Run("benchSet", func(b *testing.B) {
		b.ReportAllocs()
		for j := 0; j < b.N; j++ {
			index := j % len(zipList)
			if err := cache.Set(zipList[index], []byte("value")); err != nil {
				b.Error(err)
			}
		}
	})
	fillData(cache, zipList)
	errCount := 0
	b.Run("benchGet", func(b *testing.B) {
		b.ReportAllocs()
		for j := 0; j < b.N; j++ {
			index := j % len(zipList)
			val, err := cache.Get(zipList[index])
			if err != nil {
				errCount++
			} else {
				if !reflect.DeepEqual([]byte("value"), val) {
					b.Error("noe equal val", index)
				}
			}
		}
	})
	assert.Equal(b, cache.GetMiss(), errCount)
}

func BenchmarkGetSetParallel(b *testing.B) {
	cache := NewCache(workloadSize)
	zipList := zipfKeyList()
	b.Run("benchSet", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		index := rand.Int() % len(zipList)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				cache.Set(zipList[index], []byte("value"))
			}
		})
	})
	var errCount int32

	b.Run("benchGet", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		index := rand.Int() % len(zipList)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				val, err := cache.Get(zipList[index])
				if err != nil {
					atomic.AddInt32(&errCount, 1)
				} else {
					assert.Equal(b, []byte("value"), val)
				}
			}
		})
		assert.Equal(b, cache.GetMiss(), int(errCount))
	})
}

func TestGetSet(t *testing.T) {
	cache := NewCache(workloadSize)
	zipList := zipfKeyList()
	t.Run("benchSet", func(t *testing.T) {
		for j := 0; j < len(zipList); j++ {
			if err := cache.Set(zipList[j], []byte("value")); err != nil {
				t.Error(err)
			}
		}
	})
	fillData(cache, zipList)
	t.Run("benchGet", func(b *testing.T) {
		errCount := 0
		for j := 0; j < len(zipList); j++ {
			index := j % len(zipList)
			val, err := cache.Get(zipList[index])
			if err != nil {
				errCount++
			} else {
				if !reflect.DeepEqual([]byte("value"), val) {
					b.Error("noe equal val", index)
				}
			}
		}
		assert.Equal(b, cache.GetMiss(), errCount)
	})
}

func fillData(cache *Cache, keys [][]byte) {

	for i := 0; i < len(keys); i++ {
		cache.Set(keys[i], []byte("value"))
	}
}

func zipfKeyList() [][]byte {
	// To ensure repetition of keys in the array,
	// we are generating keys in the range from 0 to workloadSize/3.
	maxKey := int64(workloadSize) / 3

	// scrambled zipfian to ensure same keys are not together
	z := generator.NewScrambledZipfian(0, maxKey, generator.ZipfianConstant)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	keys := make([][]byte, workloadSize)
	for i := 0; i < workloadSize; i++ {
		keys[i] = []byte(strconv.Itoa(int(z.Next(r))))
	}

	return keys
}
