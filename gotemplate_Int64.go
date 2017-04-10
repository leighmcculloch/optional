package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int64 []optionalInt64

func EmptyInt64() Int64 {
	return Int64{(*emptyInt64)(nil)}
}

func OfInt64(value int64) Int64 {
	return Int64{presentInt64(value)}
}

func OfInt64Ptr(ptr *int64) Int64 {
	if ptr == nil {
		return EmptyInt64()
	} else {
		return OfInt64(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Int64) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Int64) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Int64) If(f func(value int64)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Int64) ElseFunc(f func() int64) (value int64) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int64) Else(elseValue int64) (value int64) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Int64) String() string {
	return o[0].String()
}

type optionalInt64 interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value int64))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() int64) (value int64)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue int64) (value int64)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyInt64 struct{}

func (e *emptyInt64) IsEmpty() bool {
	return true
}

func (e *emptyInt64) IsPresent() bool {
	return false
}

func (e *emptyInt64) If(f func(value int64)) {
}

func (e *emptyInt64) ElseFunc(f func() int64) (value int64) {
	return f()
}

func (e *emptyInt64) Else(elseValue int64) (value int64) {
	return elseValue
}

func (e *emptyInt64) String() string {
	return ""
}

type presentInt64 int64

func (_ presentInt64) IsEmpty() bool {
	return false
}

func (_ presentInt64) IsPresent() bool {
	return true
}

func (p presentInt64) If(f func(value int64)) {
	f(int64(p))
}

func (p presentInt64) ElseFunc(f func() int64) (value int64) {
	return int64(p)
}

func (p presentInt64) Else(elseValue int64) (value int64) {
	return int64(p)
}

func (p presentInt64) String() string {
	return fmt.Sprintf("%v", int64(p))
}
