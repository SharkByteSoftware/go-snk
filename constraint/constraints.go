// Package constraint provides useful type constraints.
package constraint

// Signed is a constraint that permits any signed integer type.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integers is a constraint that permits any integer type.
type Integers interface {
	Signed | Unsigned
}

// Floats is a constraint that permits any floating-point type.
type Floats interface {
	~float32 | ~float64
}

// Complex is a constraint that permits any complex numeric type.
type Complex interface {
	~complex64 | ~complex128
}

// Numeric is a constraint that permits any numeric type except complex.
type Numeric interface {
	Integers | Floats
}
