package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Float64 = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Float64 optionalFloat64

type optionalFloat64 []float64

const (
	valueKeyFloat64 = iota
)

// Of wraps the value in an optional.
func OfFloat64(value float64) Float64 {
	return Float64{valueKeyFloat64: value}
}

func OfFloat64Ptr(ptr *float64) Float64 {
	if ptr == nil {
		return EmptyFloat64()
	} else {
		return OfFloat64(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyFloat64() Float64 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Float64) Get() (value float64, ok bool) {
	o.If(func(v float64) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Float64) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Float64) If(f func(value float64)) {
	if o.IsPresent() {
		f(o[valueKeyFloat64])
	}
}

func (o Float64) ElseFunc(f func() float64) (value float64) {
	if o.IsPresent() {
		o.If(func(v float64) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Float64) Else(elseValue float64) (value float64) {
	return o.ElseFunc(func() float64 { return elseValue })
}

// ElseZero returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Float64) ElseZero() (value float64) {
	var zero float64
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Float64) String() string {
	return fmt.Sprintf("%v", o.ElseZero())
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Float64) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Float64) UnmarshalJSON(data []byte) error {
	var v float64
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfFloat64(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Float64) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Float64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v float64
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfFloat64(v)
	return nil
}
