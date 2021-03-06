// Code generated by gotemplate. DO NOT EDIT.

package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Int8 = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int8 optionalInt8

type optionalInt8 []int8

const (
	valueKeyInt8 = iota
)

// Of wraps the value in an optional.
func OfInt8(value int8) Int8 {
	return Int8{valueKeyInt8: value}
}

func OfInt8Ptr(ptr *int8) Int8 {
	if ptr == nil {
		return EmptyInt8()
	} else {
		return OfInt8(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyInt8() Int8 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Int8) Get() (value int8, ok bool) {
	o.If(func(v int8) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Int8) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Int8) If(f func(value int8)) {
	if o.IsPresent() {
		f(o[valueKeyInt8])
	}
}

func (o Int8) ElseFunc(f func() int8) (value int8) {
	if o.IsPresent() {
		o.If(func(v int8) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Int8) Else(elseValue int8) (value int8) {
	return o.ElseFunc(func() int8 { return elseValue })
}

// ElseZero returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Int8) ElseZero() (value int8) {
	var zero int8
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Int8) String() string {
	return fmt.Sprintf("%v", o.ElseZero())
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Int8) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Int8) UnmarshalJSON(data []byte) error {
	var v int8
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfInt8(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Int8) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Int8) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v int8
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfInt8(v)
	return nil
}
