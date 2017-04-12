package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Complex128 = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Complex128 optionalComplex128

type optionalComplex128 []complex128

const (
	valueKeyComplex128 = iota
)

// Of wraps the value in an optional.
func OfComplex128(value complex128) Complex128 {
	return Complex128{valueKeyComplex128: value}
}

func OfComplex128Ptr(ptr *complex128) Complex128 {
	if ptr == nil {
		return EmptyComplex128()
	} else {
		return OfComplex128(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyComplex128() Complex128 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Complex128) Get() (value complex128, ok bool) {
	o.If(func(v complex128) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Complex128) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Complex128) If(f func(value complex128)) {
	if o.IsPresent() {
		f(o[valueKeyComplex128])
	}
}

func (o Complex128) ElseFunc(f func() complex128) (value complex128) {
	if o.IsPresent() {
		o.If(func(v complex128) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Complex128) Else(elseValue complex128) (value complex128) {
	return o.ElseFunc(func() complex128 { return elseValue })
}

// ElseZero returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Complex128) ElseZero() (value complex128) {
	var zero complex128
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Complex128) String() string {
	return fmt.Sprintf("%v", o.ElseZero())
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Complex128) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Complex128) UnmarshalJSON(data []byte) error {
	var v complex128
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfComplex128(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Complex128) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Complex128) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v complex128
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfComplex128(v)
	return nil
}
