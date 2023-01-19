package trng

import (
	"math/rand"

	. "github.com/klauspost/cpuid/v2"
)

const (
	SUGGESTED_RETRIES = 10
)

var (
	hasRDRAND = false
)

type Rdrand struct {
	retries int
}

func init() {
	hasRDRAND = CPU.Supports(RDRAND, RDSEED)
}

func New(retries ...int) rand.Source64 {
	r := SUGGESTED_RETRIES
	if len(retries) > 0 {
		r = retries[0]
	}
	return &Rdrand{
		retries: r,
	}
}

func Uint64(ptr *uint64)

// rdrand dones't require seeding.
// only for implement.
func (r *Rdrand) Seed(seed int64) {

}

func (r *Rdrand) Int63() int64 {
	return 0
}
func (r *Rdrand) Uint64() uint64 {
	return 0
}
