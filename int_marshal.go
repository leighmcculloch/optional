package optional

import (
	"strconv"
)

// MarshalText returns text for marshaling this Int.
func (o Int) MarshalText() (text []byte, err error) {
	o.If(func(i int) {
		text = []byte(strconv.FormatInt(int64(i), 10))
	})
	return
}

// UnmarshalText parse the text into this Int
func (o *Int) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, 0)
	if err != nil {
		return err
	}
	*o = OfInt(int(v))
	return nil
}
