package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint16 []optionalUint16

func EmptyUint16() Uint16 {
	return Uint16{(*emptyUint16)(nil)}
}

func OfUint16(value uint16) Uint16 {
	return Uint16{presentUint16(value)}
}

func OfUint16Ptr(ptr *uint16) Uint16 {
	if ptr == nil {
		return EmptyUint16()
	} else {
		return OfUint16(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Uint16) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Uint16) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Uint16) If(f func(value uint16)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Uint16) ElseFunc(f func() uint16) (value uint16) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint16) Else(elseValue uint16) (value uint16) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Uint16) String() string {
	return o[0].String()
}

type optionalUint16 interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value uint16))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() uint16) (value uint16)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue uint16) (value uint16)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyUint16 struct{}

func (e *emptyUint16) IsEmpty() bool {
	return true
}

func (e *emptyUint16) IsPresent() bool {
	return false
}

func (e *emptyUint16) If(f func(value uint16)) {
}

func (e *emptyUint16) ElseFunc(f func() uint16) (value uint16) {
	return f()
}

func (e *emptyUint16) Else(elseValue uint16) (value uint16) {
	return elseValue
}

func (e *emptyUint16) String() string {
	return ""
}

type presentUint16 uint16

func (_ presentUint16) IsEmpty() bool {
	return false
}

func (_ presentUint16) IsPresent() bool {
	return true
}

func (p presentUint16) If(f func(value uint16)) {
	f(uint16(p))
}

func (p presentUint16) ElseFunc(f func() uint16) (value uint16) {
	return uint16(p)
}

func (p presentUint16) Else(elseValue uint16) (value uint16) {
	return uint16(p)
}

func (p presentUint16) String() string {
	return fmt.Sprintf("%v", uint16(p))
}
