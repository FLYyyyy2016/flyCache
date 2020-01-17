package flyCache

import (
	"github.com/FLYyyyy2016/flyCache/ringbuf"
	"hash/fnv"
	"log"
)

const segmentCount = 1024
const minSize = 1024 * 1024

type Cache struct {
	rbs      [segmentCount]*ringbuf.RingBuf
	indexMap [segmentCount]map[string]*entity
	size     int
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
	return &Cache{rbs: rbs, size: size, indexMap: indexMap}
}

func (cache *Cache) Set(key, value []byte) {
	index := getHash(key) % segmentCount
	rb := cache.rbs[index]
	m := cache.indexMap[index]
	m[string(key)] = newEntity(key, value, rb.WriteIndex())
	rb.Write(toByte(key, value))
}

func (cache *Cache) Get(key []byte) []byte {
	index := getHash(key) % segmentCount
	rb := cache.rbs[index]
	m := cache.indexMap[index]
	if e, ok := m[string(key)]; ok {
		if e.index < rb.WriteIndex()-rb.GetSize() {
			return nil
		}
		data := make([]byte, e.valueSize)
		rb.ReadAt(data, e.index+e.keySize)
		return data
	} else {
		return nil
	}
}

type entity struct {
	index, keySize, valueSize int
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
