package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Uint64 = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint64 optionalUint64

type optionalUint64 []uint64

const (
	valueKeyUint64 = iota
)

// Of wraps the value in an optional.
func OfUint64(value uint64) Uint64 {
	return Uint64{valueKeyUint64: value}
}

func OfUint64Ptr(ptr *uint64) Uint64 {
	if ptr == nil {
		return EmptyUint64()
	} else {
		return OfUint64(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyUint64() Uint64 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Uint64) Get() (value uint64, ok bool) {
	o.If(func(v uint64) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Uint64) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Uint64) If(f func(value uint64)) {
	if o.IsPresent() {
		f(o[valueKeyUint64])
	}
}

func (o Uint64) ElseFunc(f func() uint64) (value uint64) {
	if o.IsPresent() {
		o.If(func(v uint64) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Uint64) Else(elseValue uint64) (value uint64) {
	return o.ElseFunc(func() uint64 { return elseValue })
}

// ElseZero returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Uint64) ElseZero() (value uint64) {
	var zero uint64
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Uint64) String() string {
	return fmt.Sprintf("%v", o.ElseZero())
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint64) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Uint64) UnmarshalJSON(data []byte) error {
	var v uint64
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfUint64(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint64) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Uint64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v uint64
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfUint64(v)
	return nil
}
