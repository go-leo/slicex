package slicex

func Concat[S ~[]E, E comparable](ss ...S) S {
	var length int
	for _, s := range ss {
		length += len(s)
	}
	r := make(S, 0, length)
	for _, s := range ss {
		r = append(r, s...)
	}
	return r
}
