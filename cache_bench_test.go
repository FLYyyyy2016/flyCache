package flyCache

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

const (
	// based on 21million dataset, we observed a maximum key length of 77,
	// with minimum length being 6 and average length being 25. We also
	// observed that 99% of keys had length <64 bytes.
	maxKeyLength = 128
	// workloadSize is the size of array storing sequence of keys that we
	// have in our workload. In the benchmark, we iterate over this array b.N
	// number of times in circular fashion starting at a random position.
	workloadSize = 2 << 20
)


func zipfKeyList()[][]byte{
	zipf:=rand.NewZipf(rand.New(rand.NewSource(time.Now().UnixNano())),1.5,10,workloadSize/3)
	keys := make([][]byte, workloadSize)

	for i:=0;i<workloadSize;i++{
		temp:=zipf.Uint64()
		keys[i]=[]byte(strconv.Itoa(int(temp)))
	}
	return keys
}


func oneKeyList() [][]byte {
	v := rand.Int() % (workloadSize / 3)
	s := []byte(strconv.Itoa(v))

	keys := make([][]byte, workloadSize)
	for i := 0; i < workloadSize; i++ {
		keys[i] = s
	}

	return keys
}


func TestGetZipList(t *testing.T){
	zipfList:=zipfKeyList()
	oneList:=oneKeyList()
	log.Println(len(zipfList),len(oneList))
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}