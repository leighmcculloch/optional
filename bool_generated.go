package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Bool = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Bool optionalBool

type optionalBool []bool

const (
	valueKeyBool = iota
)

// Of wraps the value in an optional.
func OfBool(value bool) Bool {
	return Bool{valueKeyBool: value}
}

func OfBoolPtr(ptr *bool) Bool {
	if ptr == nil {
		return EmptyBool()
	} else {
		return OfBool(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyBool() Bool {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Bool) Get() (value bool, ok bool) {
	o.If(func(v bool) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Bool) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Bool) If(f func(value bool)) {
	if o.IsPresent() {
		f(o[valueKeyBool])
	}
}

func (o Bool) ElseFunc(f func() bool) (value bool) {
	if o.IsPresent() {
		o.If(func(v bool) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Bool) Else(elseValue bool) (value bool) {
	return o.ElseFunc(func() bool { return elseValue })
}

// ElseZero returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Bool) ElseZero() (value bool) {
	var zero bool
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Bool) String() string {
	return fmt.Sprintf("%v", o.ElseZero())
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Bool) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Bool) UnmarshalJSON(data []byte) error {
	var v bool
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfBool(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Bool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Bool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v bool
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfBool(v)
	return nil
}
