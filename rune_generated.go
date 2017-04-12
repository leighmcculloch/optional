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

// Of wraps the value in an Optional.
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

// Empty returns an empty Optional.
func EmptyRune() Rune {
	return nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Rune) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this Optional.
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

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Rune) Else(elseValue rune) (value rune) {
	return o.ElseFunc(func() rune { return elseValue })
}

// ElseZero returns the value wrapped by this Optional, or the zero value of
// the type wrapped if there is no value wrapped by this Optional.
func (o Rune) ElseZero() (value rune) {
	var zero rune
	return o.Else(zero)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Rune) String() string {
	if o.IsPresent() {
		var value rune
		o.If(func(v rune) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}

func (o Rune) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

func (o *Rune) UnmarshalJSON(data []byte) error {
	var v rune
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfRune(v)
	return nil
}

func (o Rune) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

func (o *Rune) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v rune
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfRune(v)
	return nil
}
