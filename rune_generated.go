package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Rune = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Rune optionalRune

type optionalRune []rune

const (
	valueKeyRune = iota
)

// Of wraps the value in an optional.
func OfRune(value rune) Rune {
	return Rune{valueKeyRune: value}
}

func OfRunePtr(ptr *rune) Rune {
	if ptr == nil {
		return EmptyRune()
	} else {
		return OfRune(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyRune() Rune {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Rune) Get() (value rune, ok bool) {
	o.If(func(v rune) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Rune) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Rune) If(f func(value rune)) {
	if o.IsPresent() {
		f(o[valueKeyRune])
	}
}

func (o Rune) ElseFunc(f func() rune) (value rune) {
	if o.IsPresent() {
		o.If(func(v rune) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Rune) Else(elseValue rune) (value rune) {
	return o.ElseFunc(func() rune { return elseValue })
}

// ElseZero returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Rune) ElseZero() (value rune) {
	var zero rune
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Rune) String() string {
	return fmt.Sprintf("%v", o.ElseZero())
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Rune) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Rune) UnmarshalJSON(data []byte) error {
	var v rune
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfRune(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Rune) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Rune) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v rune
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfRune(v)
	return nil
}
