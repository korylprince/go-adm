package schema

// NewValue is a helper function to get a pointer to a literal.
// e.g. NewValue[int64](10)
func NewValue[T any](v T) *T {
	return &v
}

// FromValue is a helper to get pointer values. If the pointer is nil, a zero value will be returned instead.
// e.g. FromValue[int64](NewValue[int64](10)) == 10 and FromValue[int64](nil) == 0
func FromValue[T any](v *T) T {
	if v == nil {
		var n T
		return n
	}
	return *v
}
