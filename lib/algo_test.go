package lib

import (
	"math/rand"
	"testing"
	"time"
	"sort"
)

const (
	All = 10000
	Find = 5270
	)


func TestBinarySearch(t *testing.T) {
	const MAX = 100
	rand.Seed(time.Now().Unix())
	all := rand.Intn(MAX)
	find := rand.Intn(all - 1)
	in := GenerateSlice(all)
	r := BinarySearch(find, in)
	if r == 0 {
		t.Errorf("Not found %d in %d", find, all)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sl := GenerateSlice(All)
		BinarySearch(Find, sl)
	}
}
func BenchmarkBinarySearch2(b *testing.B) {
	for i:=0; i < b.N; i++{
		sl := GenerateSlice(All)
		sort.Search(All, func(z int) bool {
			return sl[z] >= Find
		})
	}

}

func GenerateSlice(num int) []int {
	in := make([]int, 0)
	for i := 0; i < num; i++ {
		in = append(in, i)
	}
	return in
}
