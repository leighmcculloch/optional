package template

import "testing"

func TestIsPresent(t *testing.T) {
	s := "ptr to string"
	tests := []struct {
		Optional          Optional
		ExpectedIsPresent bool
	}{
		{Empty(), false},
		{Of(""), true},
		{Of("string"), true},
		{OfOptionalPtr((*T)(nil)), false},
		{OfOptionalPtr((*T)(&s)), true},
	}

	for _, test := range tests {
		isPresent := test.Optional.IsPresent()

		if isPresent != test.ExpectedIsPresent {
			t.Errorf("%#v IsPresent got %#v, want %#v", test.Optional, isPresent, test.ExpectedIsPresent)
		}
	}
}

func TestIfPresent(t *testing.T) {
	s := "ptr to string"
	tests := []struct {
		Optional       Optional
		ExpectedCalled bool
		IfCalledValue  T
	}{
		{Empty(), false, ""},
		{Of(""), true, ""},
		{Of("string"), true, "string"},
		{OfOptionalPtr((*T)(nil)), false, ""},
		{OfOptionalPtr((*T)(&s)), true, "ptr to string"},
	}

	for _, test := range tests {
		called := false
		test.Optional.If(func(v T) {
			called = true
			if v != test.IfCalledValue {
				t.Errorf("%#v IfPresent got %#v, want #%v", test.Optional, v, test.IfCalledValue)
			}
		})

		if called != test.ExpectedCalled {
			t.Errorf("%#v IfPresent called %#v, want %#v", test.Optional, called, test.ExpectedCalled)
		}
	}
}

func TestElse(t *testing.T) {
	s := "ptr to string"
	const orElse = "orelse"
	tests := []struct {
		Optional       Optional
		ExpectedResult T
	}{
		{Empty(), orElse},
		{Of(""), ""},
		{Of("string"), "string"},
		{OfOptionalPtr((*T)(nil)), orElse},
		{OfOptionalPtr((*T)(&s)), "ptr to string"},
	}

	for _, test := range tests {
		result := test.Optional.Else(orElse)

		if result != test.ExpectedResult {
			t.Errorf("%#v OrElse(%#v) got %#v, want %#v", test.Optional, orElse, result, test.ExpectedResult)
		}
	}
}

func TestElseFunc(t *testing.T) {
	s := "ptr to string"
	const orElse = "orelse"
	tests := []struct {
		Optional       Optional
		ExpectedResult T
	}{
		{Empty(), orElse},
		{Of(""), ""},
		{Of("string"), "string"},
		{OfOptionalPtr((*T)(nil)), orElse},
		{OfOptionalPtr((*T)(&s)), "ptr to string"},
	}

	for _, test := range tests {
		result := test.Optional.ElseFunc(func() T { return orElse })

		if result != test.ExpectedResult {
			t.Errorf("%#v OrElse(%#v) got %#v, want %#v", test.Optional, orElse, result, test.ExpectedResult)
		}
	}
}

func TestElseZero(t *testing.T) {
	s := "ptr to string"
	tests := []struct {
		Optional       Optional
		ExpectedResult T
	}{
		{Empty(), ""},
		{Of(""), ""},
		{Of("string"), "string"},
		{OfOptionalPtr((*T)(nil)), ""},
		{OfOptionalPtr((*T)(&s)), "ptr to string"},
	}

	for _, test := range tests {
		result := test.Optional.ElseZero()

		if result != test.ExpectedResult {
			t.Errorf("%#v ElseZero() got %#v, want %#v", test.Optional, result, test.ExpectedResult)
		}
	}
}
