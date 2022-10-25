package utils

import (
	"fmt"
	"testing"
)

func TestMT64(t *testing.T) {
	mt := NewMT19937_64()
	mt.Seed(1)
	mt.Seed(mt.Int64())
	fmt.Println(mt.Int64())
}
