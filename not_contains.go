package slicex

import "golang.org/x/exp/slices"

// NotContains reports whether v is not present in s.
func NotContains[E comparable](s []E, v E) bool {
	return !slices.Contains(s, v)
}
