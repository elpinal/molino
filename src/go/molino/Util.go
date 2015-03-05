package molino

import (
	"fmt"
)

type util struct{}

var Util util = util{}

type IHashCode interface {
	hashCode() int
}

type hashString string

//var randSeed = 1

func (h hashString) hashCode() int {
	hashcode := 0
	for i := 0; i < len(h); i++ {
		hashcode = hashcode*31 + int(h[i])
	}
	return hashcode
}

func (_ util) hash(o interface{}) int {
	if o == nil || o == "" {
		return 0
	}
	switch o.(type) {
	case string:
		return hashString(o.(string)).hashCode()
	case IHashCode:
		return o.(IHashCode).hashCode()
	}
	panic(fmt.Sprintf("Cannot create hash from %T", o))
}

func (_ util) hasheq(o interface{}) int {
	if o == nil {
		return 0
	}
	switch o.(type) {
	case IHashEq:
		return o.(IHashEq).hasheq()
	case string:
		return hashString(o.(string)).hashCode()
	case int:
		return hashInt(o.(int))
	case int64:
		return hashInt(int(o.(int64)))
	case IHashCode:
		return o.(IHashCode).hashCode()
	}
	panic(fmt.Sprintf("Cannot create hash from %T", o))
}

func (_ util) hashCombine(s, hash int) int {
	s ^= hash + 0x9e3779 + (seed << 6) + (seed >> 2)
	return s
}

func (_ util) ret1(ret ISeq, _ interface{}) ISeq {
	return ret
}

/*
func (_ util) random() int {
	a := 16807
	m := 2147483647

	lo := a * (randSeed & 0xffff)
	hi := a * (randSeed >> 16)
	lo += (hi & 0x7fff) << 16

	if lo > m {
		lo &= m
		lo++
	}
	lo += hi >> 15

	if lo > m {
		lo &= m
		lo++
	}

	randSeed = lo
	return lo
}
*/
