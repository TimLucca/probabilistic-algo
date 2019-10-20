package randsearch

import (
	"math/rand"
	"time"
)

func makeArray(size int) *[]int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := make([]int64, size)
	for i := 0; i < size; i++ {
		array[i] = r.Int63()
	}
	return &array
}

func RandomSearch() (bool, int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := makeArray(1000)
	val := (*array)[r.Int31n(1000)]
	for i := 0; i < 5000; i++ {
		if val == (*array)[r.Int31n(1000)] {
			return true, i
		}
	}
	return false, 5000
}
