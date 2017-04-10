package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uintptr []optionalUintptr

func EmptyUintptr() Uintptr {
	return Uintptr{(*emptyUintptr)(nil)}
}

func OfUintptr(value uintptr) Uintptr {
	return Uintptr{presentUintptr(value)}
}

func OfUintptrPtr(ptr *uintptr) Uintptr {
	if ptr == nil {
		return EmptyUintptr()
	} else {
		return OfUintptr(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Uintptr) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Uintptr) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Uintptr) If(f func(value uintptr)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Uintptr) ElseFunc(f func() uintptr) (value uintptr) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uintptr) Else(elseValue uintptr) (value uintptr) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Uintptr) String() string {
	return o[0].String()
}

type optionalUintptr interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value uintptr))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() uintptr) (value uintptr)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue uintptr) (value uintptr)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyUintptr struct{}

func (e *emptyUintptr) IsEmpty() bool {
	return true
}

func (e *emptyUintptr) IsPresent() bool {
	return false
}

func (e *emptyUintptr) If(f func(value uintptr)) {
}

func (e *emptyUintptr) ElseFunc(f func() uintptr) (value uintptr) {
	return f()
}

func (e *emptyUintptr) Else(elseValue uintptr) (value uintptr) {
	return elseValue
}

func (e *emptyUintptr) String() string {
	return ""
}

type presentUintptr uintptr

func (_ presentUintptr) IsEmpty() bool {
	return false
}

func (_ presentUintptr) IsPresent() bool {
	return true
}

func (p presentUintptr) If(f func(value uintptr)) {
	f(uintptr(p))
}

func (p presentUintptr) ElseFunc(f func() uintptr) (value uintptr) {
	return uintptr(p)
}

func (p presentUintptr) Else(elseValue uintptr) (value uintptr) {
	return uintptr(p)
}

func (p presentUintptr) String() string {
	return fmt.Sprintf("%v", uintptr(p))
}
