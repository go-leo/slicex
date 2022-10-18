package slicex

// IsEmpty Checks if an slice is nil or length equals 0
func IsEmpty[E any, S ~[]E](s S) bool {
	return len(s) <= 0
}
