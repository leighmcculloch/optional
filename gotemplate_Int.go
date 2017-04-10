package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int []optionalInt

func EmptyInt() Int {
	return Int{(*emptyInt)(nil)}
}

func OfInt(value int) Int {
	return Int{presentInt(value)}
}

func OfIntPtr(ptr *int) Int {
	if ptr == nil {
		return EmptyInt()
	} else {
		return OfInt(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Int) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Int) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Int) If(f func(value int)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Int) ElseFunc(f func() int) (value int) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int) Else(elseValue int) (value int) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Int) String() string {
	return o[0].String()
}

type optionalInt interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value int))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() int) (value int)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue int) (value int)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyInt struct{}

func (e *emptyInt) IsEmpty() bool {
	return true
}

func (e *emptyInt) IsPresent() bool {
	return false
}

func (e *emptyInt) If(f func(value int)) {
}

func (e *emptyInt) ElseFunc(f func() int) (value int) {
	return f()
}

func (e *emptyInt) Else(elseValue int) (value int) {
	return elseValue
}

func (e *emptyInt) String() string {
	return ""
}

type presentInt int

func (_ presentInt) IsEmpty() bool {
	return false
}

func (_ presentInt) IsPresent() bool {
	return true
}

func (p presentInt) If(f func(value int)) {
	f(int(p))
}

func (p presentInt) ElseFunc(f func() int) (value int) {
	return int(p)
}

func (p presentInt) Else(elseValue int) (value int) {
	return int(p)
}

func (p presentInt) String() string {
	return fmt.Sprintf("%v", int(p))
}
