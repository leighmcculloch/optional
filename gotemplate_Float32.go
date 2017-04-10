package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Float32 []optionalFloat32

func EmptyFloat32() Float32 {
	return Float32{(*emptyFloat32)(nil)}
}

func OfFloat32(value float32) Float32 {
	return Float32{presentFloat32(value)}
}

func OfFloat32Ptr(ptr *float32) Float32 {
	if ptr == nil {
		return EmptyFloat32()
	} else {
		return OfFloat32(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Float32) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Float32) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Float32) If(f func(value float32)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Float32) ElseFunc(f func() float32) (value float32) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Float32) Else(elseValue float32) (value float32) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Float32) String() string {
	return o[0].String()
}

type optionalFloat32 interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value float32))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() float32) (value float32)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue float32) (value float32)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyFloat32 struct{}

func (e *emptyFloat32) IsEmpty() bool {
	return true
}

func (e *emptyFloat32) IsPresent() bool {
	return false
}

func (e *emptyFloat32) If(f func(value float32)) {
}

func (e *emptyFloat32) ElseFunc(f func() float32) (value float32) {
	return f()
}

func (e *emptyFloat32) Else(elseValue float32) (value float32) {
	return elseValue
}

func (e *emptyFloat32) String() string {
	return ""
}

type presentFloat32 float32

func (_ presentFloat32) IsEmpty() bool {
	return false
}

func (_ presentFloat32) IsPresent() bool {
	return true
}

func (p presentFloat32) If(f func(value float32)) {
	f(float32(p))
}

func (p presentFloat32) ElseFunc(f func() float32) (value float32) {
	return float32(p)
}

func (p presentFloat32) Else(elseValue float32) (value float32) {
	return float32(p)
}

func (p presentFloat32) String() string {
	return fmt.Sprintf("%v", float32(p))
}
