package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int16 []optionalInt16

func EmptyInt16() Int16 {
	return Int16{(*emptyInt16)(nil)}
}

func OfInt16(value int16) Int16 {
	return Int16{presentInt16(value)}
}

func OfInt16Ptr(ptr *int16) Int16 {
	if ptr == nil {
		return EmptyInt16()
	} else {
		return OfInt16(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Int16) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Int16) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Int16) If(f func(value int16)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Int16) ElseFunc(f func() int16) (value int16) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int16) Else(elseValue int16) (value int16) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Int16) String() string {
	return o[0].String()
}

type optionalInt16 interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value int16))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() int16) (value int16)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue int16) (value int16)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyInt16 struct{}

func (e *emptyInt16) IsEmpty() bool {
	return true
}

func (e *emptyInt16) IsPresent() bool {
	return false
}

func (e *emptyInt16) If(f func(value int16)) {
}

func (e *emptyInt16) ElseFunc(f func() int16) (value int16) {
	return f()
}

func (e *emptyInt16) Else(elseValue int16) (value int16) {
	return elseValue
}

func (e *emptyInt16) String() string {
	return ""
}

type presentInt16 int16

func (_ presentInt16) IsEmpty() bool {
	return false
}

func (_ presentInt16) IsPresent() bool {
	return true
}

func (p presentInt16) If(f func(value int16)) {
	f(int16(p))
}

func (p presentInt16) ElseFunc(f func() int16) (value int16) {
	return int16(p)
}

func (p presentInt16) Else(elseValue int16) (value int16) {
	return int16(p)
}

func (p presentInt16) String() string {
	return fmt.Sprintf("%v", int16(p))
}
