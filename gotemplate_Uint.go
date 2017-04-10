package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint []optionalUint

func EmptyUint() Uint {
	return Uint{(*emptyUint)(nil)}
}

func OfUint(value uint) Uint {
	return Uint{presentUint(value)}
}

func OfUintPtr(ptr *uint) Uint {
	if ptr == nil {
		return EmptyUint()
	} else {
		return OfUint(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Uint) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Uint) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Uint) If(f func(value uint)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Uint) ElseFunc(f func() uint) (value uint) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint) Else(elseValue uint) (value uint) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Uint) String() string {
	return o[0].String()
}

type optionalUint interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value uint))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() uint) (value uint)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue uint) (value uint)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyUint struct{}

func (e *emptyUint) IsEmpty() bool {
	return true
}

func (e *emptyUint) IsPresent() bool {
	return false
}

func (e *emptyUint) If(f func(value uint)) {
}

func (e *emptyUint) ElseFunc(f func() uint) (value uint) {
	return f()
}

func (e *emptyUint) Else(elseValue uint) (value uint) {
	return elseValue
}

func (e *emptyUint) String() string {
	return ""
}

type presentUint uint

func (_ presentUint) IsEmpty() bool {
	return false
}

func (_ presentUint) IsPresent() bool {
	return true
}

func (p presentUint) If(f func(value uint)) {
	f(uint(p))
}

func (p presentUint) ElseFunc(f func() uint) (value uint) {
	return uint(p)
}

func (p presentUint) Else(elseValue uint) (value uint) {
	return uint(p)
}

func (p presentUint) String() string {
	return fmt.Sprintf("%v", uint(p))
}
