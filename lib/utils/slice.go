package utils

func SliceSetUint32(x []uint32) []uint32 {
	var m = make(map[uint32]struct{})
	for _, a := range x {
		m[a] = struct{}{}
	}
	var res = make([]uint32, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}
