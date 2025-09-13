package conditionals

// If is a ternary function for returning a value if a condition is true
// and another if false.
func If[T any](cond bool, ifTrue T, ifFalse T) T {
	if cond {
		return ifTrue
	}

	return ifFalse
}

// IfCall is a ternary function for calling a function if a condition is true
// and calling another if false.
func IfCall[T any](cond bool, truFunc func() T, falseFunc func() T) T {
	callee := If(cond, truFunc, falseFunc)
	return callee()
}
