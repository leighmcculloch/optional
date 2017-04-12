package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Uint = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint optionalUint

type optionalUint []uint

const (
	valueKeyUint = iota
)

// Of wraps the value in an optional.
func OfUint(value uint) Uint {
	return Uint{valueKeyUint: value}
}

func OfUintPtr(ptr *uint) Uint {
	if ptr == nil {
		return EmptyUint()
	} else {
		return OfUint(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyUint() Uint {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Uint) Get() (value uint, ok bool) {
	o.If(func(v uint) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Uint) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Uint) If(f func(value uint)) {
	if o.IsPresent() {
		f(o[valueKeyUint])
	}
}

func (o Uint) ElseFunc(f func() uint) (value uint) {
	if o.IsPresent() {
		o.If(func(v uint) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Uint) Else(elseValue uint) (value uint) {
	return o.ElseFunc(func() uint { return elseValue })
}

// ElseZero returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Uint) ElseZero() (value uint) {
	var zero uint
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Uint) String() string {
	return fmt.Sprintf("%v", o.ElseZero())
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Uint) UnmarshalJSON(data []byte) error {
	var v uint
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfUint(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Uint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v uint
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfUint(v)
	return nil
}
