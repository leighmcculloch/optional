// Code generated by gotemplate. DO NOT EDIT.

package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Float32 = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Float32 optionalFloat32

type optionalFloat32 []float32

const (
	valueKeyFloat32 = iota
)

// Of wraps the value in an optional.
func OfFloat32(value float32) Float32 {
	return Float32{valueKeyFloat32: value}
}

func OfFloat32Ptr(ptr *float32) Float32 {
	if ptr == nil {
		return EmptyFloat32()
	} else {
		return OfFloat32(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyFloat32() Float32 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Float32) Get() (value float32, ok bool) {
	o.If(func(v float32) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Float32) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Float32) If(f func(value float32)) {
	if o.IsPresent() {
		f(o[valueKeyFloat32])
	}
}

func (o Float32) ElseFunc(f func() float32) (value float32) {
	if o.IsPresent() {
		o.If(func(v float32) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Float32) Else(elseValue float32) (value float32) {
	return o.ElseFunc(func() float32 { return elseValue })
}

// ElseZero returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Float32) ElseZero() (value float32) {
	var zero float32
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Float32) String() string {
	return fmt.Sprintf("%v", o.ElseZero())
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Float32) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Float32) UnmarshalJSON(data []byte) error {
	var v float32
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfFloat32(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Float32) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Float32) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v float32
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfFloat32(v)
	return nil
}
