package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int8 []optionalInt8

func EmptyInt8() Int8 {
	return Int8{(*emptyInt8)(nil)}
}

func OfInt8(value int8) Int8 {
	return Int8{presentInt8(value)}
}

func OfInt8Ptr(ptr *int8) Int8 {
	if ptr == nil {
		return EmptyInt8()
	} else {
		return OfInt8(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Int8) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Int8) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Int8) If(f func(value int8)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Int8) ElseFunc(f func() int8) (value int8) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int8) Else(elseValue int8) (value int8) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Int8) String() string {
	return o[0].String()
}

type optionalInt8 interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value int8))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() int8) (value int8)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue int8) (value int8)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyInt8 struct{}

func (e *emptyInt8) IsEmpty() bool {
	return true
}

func (e *emptyInt8) IsPresent() bool {
	return false
}

func (e *emptyInt8) If(f func(value int8)) {
}

func (e *emptyInt8) ElseFunc(f func() int8) (value int8) {
	return f()
}

func (e *emptyInt8) Else(elseValue int8) (value int8) {
	return elseValue
}

func (e *emptyInt8) String() string {
	return ""
}

type presentInt8 int8

func (_ presentInt8) IsEmpty() bool {
	return false
}

func (_ presentInt8) IsPresent() bool {
	return true
}

func (p presentInt8) If(f func(value int8)) {
	f(int8(p))
}

func (p presentInt8) ElseFunc(f func() int8) (value int8) {
	return int8(p)
}

func (p presentInt8) Else(elseValue int8) (value int8) {
	return int8(p)
}

func (p presentInt8) String() string {
	return fmt.Sprintf("%v", int8(p))
}
