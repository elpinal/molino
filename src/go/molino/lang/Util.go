package lang

import (
	"fmt"
)

type util struct {}

var Util util = util{}

type IHashCode interface {
	hashCode() int
}

func (_ util) hash(o interface{}) int {
	if o == nil || o == "" {
		return 0
	}
	switch o.(type) {
	case string:
		s := o.(string)
		hashcode := 0
		for i := 0; i < len(s); i++ {
			hashcode = hashcode * 31 + int(s[i])
		}
		return hashcode
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
