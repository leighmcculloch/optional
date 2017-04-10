package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Rune []optionalRune

func EmptyRune() Rune {
	return Rune{(*emptyRune)(nil)}
}

func OfRune(value rune) Rune {
	return Rune{presentRune(value)}
}

func OfRunePtr(ptr *rune) Rune {
	if ptr == nil {
		return EmptyRune()
	} else {
		return OfRune(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Rune) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Rune) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Rune) If(f func(value rune)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Rune) ElseFunc(f func() rune) (value rune) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Rune) Else(elseValue rune) (value rune) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Rune) String() string {
	return o[0].String()
}

type optionalRune interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value rune))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() rune) (value rune)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue rune) (value rune)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyRune struct{}

func (e *emptyRune) IsEmpty() bool {
	return true
}

func (e *emptyRune) IsPresent() bool {
	return false
}

func (e *emptyRune) If(f func(value rune)) {
}

func (e *emptyRune) ElseFunc(f func() rune) (value rune) {
	return f()
}

func (e *emptyRune) Else(elseValue rune) (value rune) {
	return elseValue
}

func (e *emptyRune) String() string {
	return ""
}

type presentRune rune

func (_ presentRune) IsEmpty() bool {
	return false
}

func (_ presentRune) IsPresent() bool {
	return true
}

func (p presentRune) If(f func(value rune)) {
	f(rune(p))
}

func (p presentRune) ElseFunc(f func() rune) (value rune) {
	return rune(p)
}

func (p presentRune) Else(elseValue rune) (value rune) {
	return rune(p)
}

func (p presentRune) String() string {
	return fmt.Sprintf("%v", rune(p))
}
