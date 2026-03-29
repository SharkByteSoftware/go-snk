package slicex

// Pair holds two values of independent types.
type Pair[A any, B any] struct {
	Left  A
	Right B
}
