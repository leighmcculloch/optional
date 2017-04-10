package template

import (
	"fmt"
)

// template type Optional(T)
type T string

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Optional []optional

func Empty() Optional {
	return Optional{(*empty)(nil)}
}

func Of(value T) Optional {
	return Optional{present(value)}
}

func OfOptionalPtr(ptr *T) Optional {
	if ptr == nil {
		return Empty()
	} else {
		return Of(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Optional) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Optional) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Optional) If(f func(value T)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Optional) ElseFunc(f func() T) (value T) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Optional) Else(elseValue T) (value T) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Optional) String() string {
	return o[0].String()
}

type optional interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value T))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() T) (value T)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue T) (value T)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type empty struct{}

func (e *empty) IsEmpty() bool {
	return true
}

func (e *empty) IsPresent() bool {
	return false
}

func (e *empty) If(f func(value T)) {
}

func (e *empty) ElseFunc(f func() T) (value T) {
	return f()
}

func (e *empty) Else(elseValue T) (value T) {
	return elseValue
}

func (e *empty) String() string {
	return ""
}

type present T

func (_ present) IsEmpty() bool {
	return false
}

func (_ present) IsPresent() bool {
	return true
}

func (p present) If(f func(value T)) {
	f(T(p))
}

func (p present) ElseFunc(f func() T) (value T) {
	return T(p)
}

func (p present) Else(elseValue T) (value T) {
	return T(p)
}

func (p present) String() string {
	return fmt.Sprintf("%v", T(p))
}
