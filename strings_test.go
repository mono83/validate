package validate

import "testing"

func TestStringNotEmpty_Validate(t *testing.T) {
	dataSetNoError(
		t,
		func(v interface{}) error { return StringNotEmpty(v.(string)).Validate() },
		[]interface{}{"foo", "1", "0", " "},
	)
	dataSetError(
		t,
		func(v interface{}) error { return StringNotEmpty(v.(string)).Validate() },
		[]interface{}{""},
	)
}

func TestStringNotWhitespace_Validate(t *testing.T) {
	dataSetNoError(
		t,
		func(v interface{}) error { return StringNotWhitespace(v.(string)).Validate() },
		[]interface{}{"foo", "1", "0"},
	)
	dataSetError(
		t,
		func(v interface{}) error { return StringNotWhitespace(v.(string)).Validate() },
		[]interface{}{"", " ", " \t\n"},
	)
}

func TestStringAlpha_Validate(t *testing.T) {
	dataSetNoError(
		t,
		func(v interface{}) error { return StringAlpha(v.(string)).Validate() },
		[]interface{}{"aaa", "ZZZ", "абвГд"},
	)
	dataSetError(
		t,
		func(v interface{}) error { return StringAlpha(v.(string)).Validate() },
		[]interface{}{"123", "123abc"},
	)
}

func TestStringLatin_Validate(t *testing.T) {
	dataSetNoError(
		t,
		func(v interface{}) error { return StringLatin(v.(string)).Validate() },
		[]interface{}{"aaa", "ZZZ"},
	)
	dataSetError(
		t,
		func(v interface{}) error { return StringLatin(v.(string)).Validate() },
		[]interface{}{"123", "абвГд", "123abc"},
	)
}

func TestStringDigits_Validate(t *testing.T) {
	dataSetNoError(
		t,
		func(v interface{}) error { return StringDigits(v.(string)).Validate() },
		[]interface{}{"375057923434267", "18"},
	)
	dataSetError(
		t,
		func(v interface{}) error { return StringDigits(v.(string)).Validate() },
		[]interface{}{"3750 5792 3434 268", "18a"},
	)
}

func TestStringLuhn_Validate(t *testing.T) {
	dataSetNoError(
		t,
		func(v interface{}) error { return StringLuhn(v.(string)).Validate() },
		[]interface{}{"375057923434267", "18"},
	)
	dataSetError(
		t,
		func(v interface{}) error { return StringLuhn(v.(string)).Validate() },
		[]interface{}{"375057923434268", "22"},
	)
}
