// This package implements the xoroshiro128+ prng.
// It is very fast, but needs two uint64's to seed it.
// The reference for this prng is the
// http://xoroshiro.di.unimi.it/xoroshiro128plus.c
// implementation.
package xoroshiro

type Rand struct {
	s1, s2 uint64
}

func New(seed1 uint64, seed2 uint64) *Rand {
	return &Rand{seed1, seed2}
}

// Gets a random uint64.
func (r *Rand) Uint64() uint64 {
	answer, xor := (r.s1 + r.s2), (r.s1 ^ r.s2)
	r.s1 = ((r.s1 << 55) | (r.s1 >> (64 - 55))) ^ xor ^ (xor << 14)
	r.s2 = (xor << 36) | (xor >> (64 - 36))
	return answer
}

// Gets a random uint32.
func (r *Rand) Uint32() uint32 {
	return uint32(r.Uint64() & 0xFFFFFFFF)
}

// Gets a random int32 on the half-open interval
// [0,n).
func (r *Rand) Int32n(n int32) int32 {
	nextNum := r.Uint64()
	v := uint32(nextNum & 0xFFFFFFFF)
	prod := uint64(v) * uint64(n)
	low := uint32(prod)
	if low < uint32(n) {
		thresh := uint32(-n) % uint32(n)
		for low < thresh {
			v = uint32(nextNum >> 32)
			prod = uint64(v) * uint64(n)
			low = uint32(prod)
			if low < thresh {
				nextNum = r.Uint64()
				v = uint32(nextNum & 0xFFFFFFFF)
				prod = uint64(v) * uint64(n)
				low = uint32(prod)
			}
		}
	}
	return int32(prod >> 32)
}
