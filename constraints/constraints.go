// Package constraints provides type constraints.
package constraints

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integers interface {
	Signed | Unsigned
}

type Floats interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

type Numeric interface {
	Integers | Floats
}
