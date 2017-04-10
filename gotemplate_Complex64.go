package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Complex64 []optionalComplex64

func EmptyComplex64() Complex64 {
	return Complex64{(*emptyComplex64)(nil)}
}

func OfComplex64(value complex64) Complex64 {
	return Complex64{presentComplex64(value)}
}

func OfComplex64Ptr(ptr *complex64) Complex64 {
	if ptr == nil {
		return EmptyComplex64()
	} else {
		return OfComplex64(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Complex64) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Complex64) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Complex64) If(f func(value complex64)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Complex64) ElseFunc(f func() complex64) (value complex64) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Complex64) Else(elseValue complex64) (value complex64) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Complex64) String() string {
	return o[0].String()
}

type optionalComplex64 interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value complex64))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() complex64) (value complex64)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue complex64) (value complex64)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyComplex64 struct{}

func (e *emptyComplex64) IsEmpty() bool {
	return true
}

func (e *emptyComplex64) IsPresent() bool {
	return false
}

func (e *emptyComplex64) If(f func(value complex64)) {
}

func (e *emptyComplex64) ElseFunc(f func() complex64) (value complex64) {
	return f()
}

func (e *emptyComplex64) Else(elseValue complex64) (value complex64) {
	return elseValue
}

func (e *emptyComplex64) String() string {
	return ""
}

type presentComplex64 complex64

func (_ presentComplex64) IsEmpty() bool {
	return false
}

func (_ presentComplex64) IsPresent() bool {
	return true
}

func (p presentComplex64) If(f func(value complex64)) {
	f(complex64(p))
}

func (p presentComplex64) ElseFunc(f func() complex64) (value complex64) {
	return complex64(p)
}

func (p presentComplex64) Else(elseValue complex64) (value complex64) {
	return complex64(p)
}

func (p presentComplex64) String() string {
	return fmt.Sprintf("%v", complex64(p))
}
