package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Uintptr = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uintptr optionalUintptr

type optionalUintptr []uintptr

const (
	valueKeyUintptr = iota
)

// Of wraps the value in an optional.
func OfUintptr(value uintptr) Uintptr {
	return Uintptr{valueKeyUintptr: value}
}

func OfUintptrPtr(ptr *uintptr) Uintptr {
	if ptr == nil {
		return EmptyUintptr()
	} else {
		return OfUintptr(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyUintptr() Uintptr {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Uintptr) Get() (value uintptr, ok bool) {
	o.If(func(v uintptr) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Uintptr) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Uintptr) If(f func(value uintptr)) {
	if o.IsPresent() {
		f(o[valueKeyUintptr])
	}
}

func (o Uintptr) ElseFunc(f func() uintptr) (value uintptr) {
	if o.IsPresent() {
		o.If(func(v uintptr) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Uintptr) Else(elseValue uintptr) (value uintptr) {
	return o.ElseFunc(func() uintptr { return elseValue })
}

// ElseZero returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Uintptr) ElseZero() (value uintptr) {
	var zero uintptr
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Uintptr) String() string {
	return fmt.Sprintf("%v", o.ElseZero())
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uintptr) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Uintptr) UnmarshalJSON(data []byte) error {
	var v uintptr
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfUintptr(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uintptr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Uintptr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v uintptr
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfUintptr(v)
	return nil
}
