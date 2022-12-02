package slicex

import "golang.org/x/exp/slices"

// ContainsFunc reports whether v is present in s.
func ContainsFunc[E comparable](s []E, f func(E) bool) bool {
	return slices.IndexFunc(s, f) >= 0
}
