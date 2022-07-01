package main

func SumIntsorFlats[K comparable, v Number](m map[K]v) v {
	var s v
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int64 | float64
}
