package slicex

// Map 方法创建一个新数组，这个新数组由原数组中的每个元素都调用一次提供的函数后的返回值组成。
func Map[S1 ~[]E1, S2 ~[]E2, E1 any, E2 any](s S1, callbackFn func(E1) E2) S2 {
	s2 := make(S2, 0, len(s))
	for _, e1 := range s {
		s2 = append(s2, callbackFn(e1))
	}
	return s2
}
