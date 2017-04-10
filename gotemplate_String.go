package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type String []optionalString

func EmptyString() String {
	return String{(*emptyString)(nil)}
}

func OfString(value string) String {
	return String{presentString(value)}
}

func OfStringPtr(ptr *string) String {
	if ptr == nil {
		return EmptyString()
	} else {
		return OfString(*ptr)
	}
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o String) IsEmpty() bool {
	return o[0].IsEmpty()
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o String) IsPresent() bool {
	return o[0].IsPresent()
}

// If calls the function if there is a value wrapped by this Optional.
func (o String) If(f func(value string)) {
	o[0].If(f)
}

// ElseFunc calls the function if there is no value wrapped by this Optional,
// and returns the value returned by the value.
func (o String) ElseFunc(f func() string) (value string) {
	return o[0].ElseFunc(f)
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o String) Else(elseValue string) (value string) {
	return o[0].Else(elseValue)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o String) String() string {
	return o[0].String()
}

type optionalString interface {
	// IsEmpty returns true if there there is no value wrapped by this Optional.
	IsEmpty() bool
	// IsPresent returns true if there is a value wrapped by this Optional.
	IsPresent() bool
	// If calls the function if there is a value wrapped by this Optional.
	If(f func(value string))
	// ElseFunc calls the function if there is no value wrapped by this Optional,
	// and returns the value returned by the value.
	ElseFunc(f func() string) (value string)
	// Else returns the value wrapped by this Optional, or the value passed in if
	// there is no value wrapped by this Optional.
	Else(elseValue string) (value string)
	// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
	String() string
}

type emptyString struct{}

func (e *emptyString) IsEmpty() bool {
	return true
}

func (e *emptyString) IsPresent() bool {
	return false
}

func (e *emptyString) If(f func(value string)) {
}

func (e *emptyString) ElseFunc(f func() string) (value string) {
	return f()
}

func (e *emptyString) Else(elseValue string) (value string) {
	return elseValue
}

func (e *emptyString) String() string {
	return ""
}

type presentString string

func (_ presentString) IsEmpty() bool {
	return false
}

func (_ presentString) IsPresent() bool {
	return true
}

func (p presentString) If(f func(value string)) {
	f(string(p))
}

func (p presentString) ElseFunc(f func() string) (value string) {
	return string(p)
}

func (p presentString) Else(elseValue string) (value string) {
	return string(p)
}

func (p presentString) String() string {
	return fmt.Sprintf("%v", string(p))
}
