package trng

import (
	"testing"
)

func TestTrng(t *testing.T) {
	var a uint64
	Uint64(&a)
	t.Log(a)
}
