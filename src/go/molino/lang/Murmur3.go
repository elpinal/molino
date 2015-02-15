package lang

const (
	seed = 0
	C1 uint = 0xcc9e2d51
	C2 uint = 0x1b873593
)

func hashInt(input int) uint {
	if input == 0 {
		return 0
	}
	var low int = input
	var high uint = uint(input) >> 32

	var k1 uint = mixK1(uint(low))
	var h1 uint = mixH1(seed, k1)

	k1 = mixK1(high)
	h1 = mixH1(h1, k1)

	return fmix(h1, 8)
}


func mixK1(k1 uint) uint {
	k1 *= C1
	k1 = (k1 << 15) | (k1 >> (32 - 15))
	k1 *= C2
	return k1
}

func mixH1(h1, k1 uint) uint {
	h1 ^= k1
	h1 = (h1 << 13) | (h1 >> (32 - 13))
	h1 = h1 * 5 + 0xe6546b64
	return h1
}

func fmix(h1, length uint) uint {
	h1 ^= length
	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16
	return h1
}
