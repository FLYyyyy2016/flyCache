package flyCache

import (
	"github.com/FLYyyyy2016/flyCache/ringbuf"
	"hash/fnv"
	"log"
	"sync"
)

const segmentCount = 1024
const minSize = 1024 * 1024

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "not found"
}

type TooLongError struct{}

func (e TooLongError) Error() string {
	return "key or value too long"
}

type KeyRewrite struct{}

func (e KeyRewrite) Error() string {
	return "has rewrite"
}

type Cache struct {
	rbs       [segmentCount]*ringbuf.RingBuf
	indexMap  [segmentCount]map[string]*entity
	size      int
	count     int
	missTimes int
	maxLength int
	sync.RWMutex
}

func NewCache(size int) *Cache {
	if size < minSize {
		size = minSize
	}
	var rbs = [segmentCount]*ringbuf.RingBuf{}
	var indexMap = [segmentCount]map[string]*entity{}
	for i := 0; i < segmentCount; i++ {
		rbs[i] = ringbuf.New(size / 1024)
		indexMap[i] = make(map[string]*entity)
	}
	return &Cache{rbs: rbs, size: size, indexMap: indexMap, maxLength: size / segmentCount / 20}
}

func (cache *Cache) GetCount() int {
	cache.RLock()
	defer cache.RUnlock()
	return cache.count
}

func (cache *Cache) GetMiss() int {
	cache.RLock()
	defer cache.RUnlock()
	return cache.missTimes
}

func (cache *Cache) Set(key, value []byte) error {
	if len(key) > cache.maxLength || len(value) > cache.maxLength {
		return TooLongError{}
	}
	index := getHash(key) % segmentCount
	rb := cache.rbs[index]
	m := cache.indexMap[index]
	rb.Lock()
	defer rb.Unlock()
	if _, ok := m[string(key)]; !ok {
		cache.addCount()
	}
	m[string(key)] = newEntity(key, value, rb.WriteIndex())
	rb.Write(toByte(key, value))
	return nil
}

func (cache *Cache) Get(key []byte) ([]byte, error) {
	index := getHash(key) % segmentCount
	rb := cache.rbs[index]
	m := cache.indexMap[index]
	rb.RLock()
	defer rb.RUnlock()
	if e, ok := m[string(key)]; ok {
		if e.index < rb.WriteIndex()-rb.GetSize() {
			cache.missCount()
			return nil, KeyRewrite{}
		}
		data := make([]byte, e.valueSize)
		rb.ReadAt(data, e.index+e.keySize)
		return data, nil
	} else {
		cache.missCount()
		return nil, NotFoundError{}
	}
}

func (cache *Cache) addCount() {
	cache.Lock()
	defer cache.Unlock()
	cache.count++
}

func (cache *Cache) missCount() {
	cache.Lock()
	defer cache.Unlock()
	cache.missTimes++
}

type entity struct {
	index, keySize, valueSize, cap int
}

func newEntity(key, value []byte, index int) *entity {
	return &entity{
		keySize:   len(key),
		valueSize: len(value),
		index:     index,
	}
}

func toByte(key, value []byte) []byte {
	result := make([]byte, 0)
	result = append(result, key...)
	result = append(result, value...)
	return result
}

func getHash(data []byte) uint32 {
	h := fnv.New32()
	length, err := h.Write(data)
	if length != len(data) || err != nil {
		log.Fatalln("error on hash")
	}
	return h.Sum32()
}
