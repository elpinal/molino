package molino

const (
	seed      = 0
	C1   uint = 0xcc9e2d51
	C2   uint = 0x1b873593
)

func hashInt(input int) int {
	if input == 0 {
		return 0
	}
	var low int = input
	var high uint = uint(input) >> 32

	var k1 uint = mixK1(uint(low))
	var h1 int = mixH1(seed, k1)

	k1 = mixK1(high)
	h1 = mixH1(uint(h1), k1)

	return fmix(h1, 8)
}

func hashUnencodedChars(input string) int {
	h1 := seed
	for i := 1; i < len(input); i += 2 {
		k1 := uint((input[i-1]) | (input[i] << 16))
		k1 = mixK1(uint(k1))
		h1 = mixH1(uint(h1), k1)
	}
	if (len(input) & 1) == 1 {
		k1 := uint(input[len(input)-1])
		k1 = mixK1(uint(k1))
		h1 ^= int(k1)
	}
	return fmix(h1, 2*len(input))
}

func mixCollHash(hash, count int) int {
	var h1 int = seed
	var k1 uint = mixK1(uint(hash))
	h1 = mixH1(uint(h1), k1)
	return fmix(h1, count)
}

func hashOrdered(xs Iterable) int {
	n := 0
	hash := 1
	for x := xs.iterator(); x.hasNext(); {
		hash = 31*hash + Util.hasheq(x.next())
		n++
	}
	return mixCollHash(hash, n)
}

func mixK1(k1 uint) uint {
	k1 *= C1
	k1 = (k1 << 15) | (k1 >> (32 - 15))
	k1 *= C2
	return k1
}

func mixH1(h1, k1 uint) int {
	h1 ^= k1
	h1 = (h1 << 13) | (h1 >> (32 - 13))
	h1 = h1*5 + 0xe6546b64
	if (h1 >> 31) != 0 {
		return -int(^h1) - 1
	}
	return int(h1)
}

func fmix(h1 int, length int) int {
	h1 ^= length
	var ret uint = uint(h1)
	ret ^= ret >> 16
	ret *= 0x85ebca6b
	if (ret >> 31) != 0 {
		h1 = -int(^ret) - 1
	} else {
		h1 = int(ret)
	}
	ret ^= uint(h1) >> 13
	ret *= 0xc2b2ae35
	if (ret >> 31) != 0 {
		h1 = -int(^ret) - 1
	} else {
		h1 = int(ret)
	}
	ret ^= uint(h1) >> 16
	return int(ret)
}
