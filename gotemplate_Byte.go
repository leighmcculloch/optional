package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Byte []optionalByte

func EmptyByte() Byte {
	return Byte{(*emptyByte)(nil)}
}

func OfByte(value byte) Byte {
	return Byte{presentByte(value)}
}

func OfBytePtr(ptr *byte) Byte {
	if ptr == nil {
		return EmptyByte()
	} else {
		return OfByte(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Byte) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Byte) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Byte) If(f func(value byte)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o Byte) ElseFunc(f func() byte) (value byte) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Byte) Else(elseValue byte) (value byte) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Byte) String() string {
	return o[0].String()
}

type optionalByte interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value byte))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() byte) (value byte)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue byte) (value byte)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyByte struct{}

func (e *emptyByte) IsEmpty() bool {
	return true
}

func (e *emptyByte) IsPresent() bool {
	return false
}

func (e *emptyByte) If(f func(value byte)) {
}

func (e *emptyByte) ElseFunc(f func() byte) (value byte) {
	return f()
}

func (e *emptyByte) Else(elseValue byte) (value byte) {
	return elseValue
}

func (e *emptyByte) String() string {
	return ""
}

type presentByte byte

func (_ presentByte) IsEmpty() bool {
	return false
}

func (_ presentByte) IsPresent() bool {
	return true
}

func (p presentByte) If(f func(value byte)) {
	f(byte(p))
}

func (p presentByte) ElseFunc(f func() byte) (value byte) {
	return byte(p)
}

func (p presentByte) Else(elseValue byte) (value byte) {
	return byte(p)
}

func (p presentByte) String() string {
	return fmt.Sprintf("%v", byte(p))
}
