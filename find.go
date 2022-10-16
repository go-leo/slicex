package slicex

// Find 方法返回数组中满足提供的函数的第一个元素的值。
func Find[S ~[]E, E any](s S, f func(int, E) bool) E {
	var e E
	for i, v := range s {
		if f(i, v) {
			e = v
			break
		}
	}
	return e
}
