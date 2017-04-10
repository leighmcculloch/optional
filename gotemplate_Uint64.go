package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint64 []optionalUint64

func EmptyUint64() Uint64 {
	return Uint64{(*emptyUint64)(nil)}
}

func OfUint64(value uint64) Uint64 {
	return Uint64{presentUint64(value)}
}

func OfUint64Ptr(ptr *uint64) Uint64 {
	if ptr == nil {
		return EmptyUint64()
	} else {
		return OfUint64(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Uint64) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Uint64) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Uint64) If(f func(value uint64)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Uint64) ElseFunc(f func() uint64) (value uint64) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint64) Else(elseValue uint64) (value uint64) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Uint64) String() string {
	return o[0].String()
}

type optionalUint64 interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value uint64))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() uint64) (value uint64)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue uint64) (value uint64)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyUint64 struct{}

func (e *emptyUint64) IsEmpty() bool {
	return true
}

func (e *emptyUint64) IsPresent() bool {
	return false
}

func (e *emptyUint64) If(f func(value uint64)) {
}

func (e *emptyUint64) ElseFunc(f func() uint64) (value uint64) {
	return f()
}

func (e *emptyUint64) Else(elseValue uint64) (value uint64) {
	return elseValue
}

func (e *emptyUint64) String() string {
	return ""
}

type presentUint64 uint64

func (_ presentUint64) IsEmpty() bool {
	return false
}

func (_ presentUint64) IsPresent() bool {
	return true
}

func (p presentUint64) If(f func(value uint64)) {
	f(uint64(p))
}

func (p presentUint64) ElseFunc(f func() uint64) (value uint64) {
	return uint64(p)
}

func (p presentUint64) Else(elseValue uint64) (value uint64) {
	return uint64(p)
}

func (p presentUint64) String() string {
	return fmt.Sprintf("%v", uint64(p))
}
