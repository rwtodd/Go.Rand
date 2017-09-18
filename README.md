# Go.Rand

This repo is for prng impelementations.  The one in the
std lib is pretty good, but for fast simulations others
may be more appropriate.

As of now, there is only one prng here: the xoroshiro+128
prng.  Very fast.  It needs two uint64s to seed it, and
I just use the normal Go `math/rand` prng to do so:

    rand.Seed(time.Now().Unix())
    xrnd := xoroshiro.New(rand.Uint64(), rand.Uint64())

