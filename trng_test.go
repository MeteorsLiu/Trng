package trng

import (
	"testing"
)

func TestTrng(t *testing.T) {
	var a uint64
	rdrand(&a)
	t.Log(a)
}
