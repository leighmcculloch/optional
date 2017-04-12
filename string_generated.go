package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _String = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type String optionalString

type optionalString []string

const (
	valueKeyString = iota
)

// Of wraps the value in an optional.
func OfString(value string) String {
	return String{valueKeyString: value}
}

func OfStringPtr(ptr *string) String {
	if ptr == nil {
		return EmptyString()
	} else {
		return OfString(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyString() String {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o String) Get() (value string, ok bool) {
	o.If(func(v string) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o String) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o String) If(f func(value string)) {
	if o.IsPresent() {
		f(o[valueKeyString])
	}
}

func (o String) ElseFunc(f func() string) (value string) {
	if o.IsPresent() {
		o.If(func(v string) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o String) Else(elseValue string) (value string) {
	return o.ElseFunc(func() string { return elseValue })
}

// ElseZero returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o String) ElseZero() (value string) {
	var zero string
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o String) String() string {
	return fmt.Sprintf("%v", o.ElseZero())
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o String) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *String) UnmarshalJSON(data []byte) error {
	var v string
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfString(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o String) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *String) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfString(v)
	return nil
}
