package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint8 []optionalUint8

func EmptyUint8() Uint8 {
	return Uint8{(*emptyUint8)(nil)}
}

func OfUint8(value uint8) Uint8 {
	return Uint8{presentUint8(value)}
}

func OfUint8Ptr(ptr *uint8) Uint8 {
	if ptr == nil {
		return EmptyUint8()
	} else {
		return OfUint8(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Uint8) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Uint8) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Uint8) If(f func(value uint8)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Uint8) ElseFunc(f func() uint8) (value uint8) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint8) Else(elseValue uint8) (value uint8) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Uint8) String() string {
	return o[0].String()
}

type optionalUint8 interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value uint8))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() uint8) (value uint8)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue uint8) (value uint8)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyUint8 struct{}

func (e *emptyUint8) IsEmpty() bool {
	return true
}

func (e *emptyUint8) IsPresent() bool {
	return false
}

func (e *emptyUint8) If(f func(value uint8)) {
}

func (e *emptyUint8) ElseFunc(f func() uint8) (value uint8) {
	return f()
}

func (e *emptyUint8) Else(elseValue uint8) (value uint8) {
	return elseValue
}

func (e *emptyUint8) String() string {
	return ""
}

type presentUint8 uint8

func (_ presentUint8) IsEmpty() bool {
	return false
}

func (_ presentUint8) IsPresent() bool {
	return true
}

func (p presentUint8) If(f func(value uint8)) {
	f(uint8(p))
}

func (p presentUint8) ElseFunc(f func() uint8) (value uint8) {
	return uint8(p)
}

func (p presentUint8) Else(elseValue uint8) (value uint8) {
	return uint8(p)
}

func (p presentUint8) String() string {
	return fmt.Sprintf("%v", uint8(p))
}
