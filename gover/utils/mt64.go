package utils

type MT19937_64 struct {
	mt  []uint64
	mti int
}

func (c *MT19937_64) Seed(seed uint64) {
	c.mt[0] = seed & 0xffffffffffffffff
	for i := uint64(1); i < 312; i++ {
		c.mt[i] = (6364136223846793005*(c.mt[i-1]^(c.mt[i-1]>>62)) + i) & 0xffffffffffffffff
	}
	c.mti = 312
}

func (c *MT19937_64) Int64() uint64 {
	if c.mti >= 312 {
		if c.mti == 313 {
			c.Seed(5489)
		}
		for k := 0; k < 311; k++ {
			y := (c.mt[k] & 0xFFFFFFFF80000000) | (c.mt[k+1] & 0x7fffffff)
			if k < 156 {
				var r uint64
				if (y & 1) != 0 {
					r = 0xB5026F5AA96619E9
				}
				c.mt[k] = c.mt[k+156] ^ (y >> 1) ^ r
			} else {
				var r uint64
				if (y & 1) != 0 {
					r = 0xB5026F5AA96619E9
				}
				c.mt[k] = c.mt[k-156] ^ (y >> 1) ^ r
			}
		}
		yy := (c.mt[311] & 0xFFFFFFFF80000000) | (c.mt[0] & 0x7fffffff)
		var r uint64
		if (yy & 1) != 0 {
			r = 0xB5026F5AA96619E9
		}
		c.mt[311] = c.mt[155] ^ (yy >> 1) ^ r
		c.mti = 0
	}
	x := c.mt[c.mti]
	c.mti += 1
	x ^= (x >> 29) & 0x5555555555555555
	x ^= (x << 17) & 0x71D67FFFEDA60000
	x ^= (x << 37) & 0xFFF7EEE000000000
	x ^= (x >> 43)
	return x
}

func NewMT19937_64() *MT19937_64 {
	return &MT19937_64{
		mt:  make([]uint64, 312),
		mti: 313,
	}
}
