package lang

type util struct {}

var Util util = util{}

func (_ util) hash(o interface{}) int {
	if o == nil || o == "" {
		return 0
	}
	panic("Cannot create hash")
}

func (_ util) hashCombine(s, hash int) int {
	s ^= hash + 0x9e3779 + (seed << 6) + (seed >> 2)
	return s
}

func (_ util) ret1(ret ISeq, _ interface{}) ISeq {
	return ret
}
