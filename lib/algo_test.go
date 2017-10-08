package lib

import (
	"testing"
	"math/rand"

	"time"
)

func TestBinarySearch(t *testing.T) {
	const MAX  = 100
	rand.Seed(time.Now().Unix())
	all := rand.Intn(MAX)
	find := rand.Intn(all - 1)
	in := make([]int, 0)
	for i := 0; i<all; i++{
		in = append(in, i)
	}
	r := BinarySearch(find, in)
	if r == 0{
		t.Errorf("Not found %d in %d", find, all)
	}
}