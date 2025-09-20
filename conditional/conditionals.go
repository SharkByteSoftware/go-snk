package conditional

// If is a ternary function for returning a value if a condition is true
// and another if false.
func If[T any](cond bool, ifTrue T, ifFalse T) T {
	if cond {
		return ifTrue
	}

	return ifFalse
}

// IfNotNil calls a function if x is not nil.
func IfNotNil[T any](x *T, callee func()) {
	if x == nil {
		return
	}

	callee()
}

// IfCall is a ternary function for calling a function if a condition is true
// and calling another if false.
func IfCall(cond bool, trueFunc func(), falseFunc func()) {
	callee := If(cond, trueFunc, falseFunc)
	callee()
}

// IfCallReturn is a ternary function for calling a function if a condition is true
// and calling another if false and returns the result.
func IfCallReturn[T any](cond bool, trueFunc func() T, falseFunc func() T) T {
	return If(cond, trueFunc, falseFunc)()
}
