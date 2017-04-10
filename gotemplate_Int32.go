package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int32 []optionalInt32

func EmptyInt32() Int32 {
	return Int32{(*emptyInt32)(nil)}
}

func OfInt32(value int32) Int32 {
	return Int32{presentInt32(value)}
}

func OfInt32Ptr(ptr *int32) Int32 {
	if ptr == nil {
		return EmptyInt32()
	} else {
		return OfInt32(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Int32) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Int32) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Int32) If(f func(value int32)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Int32) ElseFunc(f func() int32) (value int32) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int32) Else(elseValue int32) (value int32) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Int32) String() string {
	return o[0].String()
}

type optionalInt32 interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value int32))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() int32) (value int32)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue int32) (value int32)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyInt32 struct{}

func (e *emptyInt32) IsEmpty() bool {
	return true
}

func (e *emptyInt32) IsPresent() bool {
	return false
}

func (e *emptyInt32) If(f func(value int32)) {
}

func (e *emptyInt32) ElseFunc(f func() int32) (value int32) {
	return f()
}

func (e *emptyInt32) Else(elseValue int32) (value int32) {
	return elseValue
}

func (e *emptyInt32) String() string {
	return ""
}

type presentInt32 int32

func (_ presentInt32) IsEmpty() bool {
	return false
}

func (_ presentInt32) IsPresent() bool {
	return true
}

func (p presentInt32) If(f func(value int32)) {
	f(int32(p))
}

func (p presentInt32) ElseFunc(f func() int32) (value int32) {
	return int32(p)
}

func (p presentInt32) Else(elseValue int32) (value int32) {
	return int32(p)
}

func (p presentInt32) String() string {
	return fmt.Sprintf("%v", int32(p))
}
