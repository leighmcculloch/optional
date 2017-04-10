package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Bool []optionalBool

func EmptyBool() Bool {
	return Bool{(*emptyBool)(nil)}
}

func OfBool(value bool) Bool {
	return Bool{presentBool(value)}
}

func OfBoolPtr(ptr *bool) Bool {
	if ptr == nil {
		return EmptyBool()
	} else {
		return OfBool(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Bool) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Bool) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Bool) If(f func(value bool)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Bool) ElseFunc(f func() bool) (value bool) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Bool) Else(elseValue bool) (value bool) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Bool) String() string {
	return o[0].String()
}

type optionalBool interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value bool))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() bool) (value bool)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue bool) (value bool)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyBool struct{}

func (e *emptyBool) IsEmpty() bool {
	return true
}

func (e *emptyBool) IsPresent() bool {
	return false
}

func (e *emptyBool) If(f func(value bool)) {
}

func (e *emptyBool) ElseFunc(f func() bool) (value bool) {
	return f()
}

func (e *emptyBool) Else(elseValue bool) (value bool) {
	return elseValue
}

func (e *emptyBool) String() string {
	return ""
}

type presentBool bool

func (_ presentBool) IsEmpty() bool {
	return false
}

func (_ presentBool) IsPresent() bool {
	return true
}

func (p presentBool) If(f func(value bool)) {
	f(bool(p))
}

func (p presentBool) ElseFunc(f func() bool) (value bool) {
	return bool(p)
}

func (p presentBool) Else(elseValue bool) (value bool) {
	return bool(p)
}

func (p presentBool) String() string {
	return fmt.Sprintf("%v", bool(p))
}
