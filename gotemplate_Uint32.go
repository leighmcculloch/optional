package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint32 []optionalUint32

func EmptyUint32() Uint32 {
	return Uint32{(*emptyUint32)(nil)}
}

func OfUint32(value uint32) Uint32 {
	return Uint32{presentUint32(value)}
}

func OfUint32Ptr(ptr *uint32) Uint32 {
	if ptr == nil {
		return EmptyUint32()
	} else {
		return OfUint32(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Uint32) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Uint32) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Uint32) If(f func(value uint32)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Uint32) ElseFunc(f func() uint32) (value uint32) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint32) Else(elseValue uint32) (value uint32) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Uint32) String() string {
	return o[0].String()
}

type optionalUint32 interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value uint32))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() uint32) (value uint32)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue uint32) (value uint32)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyUint32 struct{}

func (e *emptyUint32) IsEmpty() bool {
	return true
}

func (e *emptyUint32) IsPresent() bool {
	return false
}

func (e *emptyUint32) If(f func(value uint32)) {
}

func (e *emptyUint32) ElseFunc(f func() uint32) (value uint32) {
	return f()
}

func (e *emptyUint32) Else(elseValue uint32) (value uint32) {
	return elseValue
}

func (e *emptyUint32) String() string {
	return ""
}

type presentUint32 uint32

func (_ presentUint32) IsEmpty() bool {
	return false
}

func (_ presentUint32) IsPresent() bool {
	return true
}

func (p presentUint32) If(f func(value uint32)) {
	f(uint32(p))
}

func (p presentUint32) ElseFunc(f func() uint32) (value uint32) {
	return uint32(p)
}

func (p presentUint32) Else(elseValue uint32) (value uint32) {
	return uint32(p)
}

func (p presentUint32) String() string {
	return fmt.Sprintf("%v", uint32(p))
}
