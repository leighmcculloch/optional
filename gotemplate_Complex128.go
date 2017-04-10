package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Complex128 []optionalComplex128

func EmptyComplex128() Complex128 {
	return Complex128{(*emptyComplex128)(nil)}
}

func OfComplex128(value complex128) Complex128 {
	return Complex128{presentComplex128(value)}
}

func OfComplex128Ptr(ptr *complex128) Complex128 {
	if ptr == nil {
		return EmptyComplex128()
	} else {
		return OfComplex128(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Complex128) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Complex128) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Complex128) If(f func(value complex128)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Complex128) ElseFunc(f func() complex128) (value complex128) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Complex128) Else(elseValue complex128) (value complex128) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Complex128) String() string {
	return o[0].String()
}

type optionalComplex128 interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value complex128))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() complex128) (value complex128)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue complex128) (value complex128)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyComplex128 struct{}

func (e *emptyComplex128) IsEmpty() bool {
	return true
}

func (e *emptyComplex128) IsPresent() bool {
	return false
}

func (e *emptyComplex128) If(f func(value complex128)) {
}

func (e *emptyComplex128) ElseFunc(f func() complex128) (value complex128) {
	return f()
}

func (e *emptyComplex128) Else(elseValue complex128) (value complex128) {
	return elseValue
}

func (e *emptyComplex128) String() string {
	return ""
}

type presentComplex128 complex128

func (_ presentComplex128) IsEmpty() bool {
	return false
}

func (_ presentComplex128) IsPresent() bool {
	return true
}

func (p presentComplex128) If(f func(value complex128)) {
	f(complex128(p))
}

func (p presentComplex128) ElseFunc(f func() complex128) (value complex128) {
	return complex128(p)
}

func (p presentComplex128) Else(elseValue complex128) (value complex128) {
	return complex128(p)
}

func (p presentComplex128) String() string {
	return fmt.Sprintf("%v", complex128(p))
}
