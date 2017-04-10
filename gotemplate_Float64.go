package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Float64 []optionalFloat64

func EmptyFloat64() Float64 {
	return Float64{(*emptyFloat64)(nil)}
}

func OfFloat64(value float64) Float64 {
	return Float64{presentFloat64(value)}
}

func OfFloat64Ptr(ptr *float64) Float64 {
	if ptr == nil {
		return EmptyFloat64()
	} else {
		return OfFloat64(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Float64) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Float64) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Float64) If(f func(value float64)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Float64) ElseFunc(f func() float64) (value float64) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Float64) Else(elseValue float64) (value float64) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Float64) String() string {
	return o[0].String()
}

type optionalFloat64 interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value float64))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() float64) (value float64)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue float64) (value float64)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyFloat64 struct{}

func (e *emptyFloat64) IsEmpty() bool {
	return true
}

func (e *emptyFloat64) IsPresent() bool {
	return false
}

func (e *emptyFloat64) If(f func(value float64)) {
}

func (e *emptyFloat64) ElseFunc(f func() float64) (value float64) {
	return f()
}

func (e *emptyFloat64) Else(elseValue float64) (value float64) {
	return elseValue
}

func (e *emptyFloat64) String() string {
	return ""
}

type presentFloat64 float64

func (_ presentFloat64) IsEmpty() bool {
	return false
}

func (_ presentFloat64) IsPresent() bool {
	return true
}

func (p presentFloat64) If(f func(value float64)) {
	f(float64(p))
}

func (p presentFloat64) ElseFunc(f func() float64) (value float64) {
	return float64(p)
}

func (p presentFloat64) Else(elseValue float64) (value float64) {
	return float64(p)
}

func (p presentFloat64) String() string {
	return fmt.Sprintf("%v", float64(p))
}
