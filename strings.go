package validate

import "strings"

// StringNotEmpty is empty string validator
type StringNotEmpty string

// Validate performs validation
func (s StringNotEmpty) Validate() error {
	if len(s) == 0 {
		return errorf(s, "Non empty string expected")
	}

	return nil
}

// StringNotWhitespace is validation for empty and whitespace strings
type StringNotWhitespace string

// Validate performs validation
func (s StringNotWhitespace) Validate() error {
	if len(s) == 0 {
		return errorf(s, "Non empty string expected")
	}
	if len(strings.TrimSpace(string(s))) == 0 {
		return errorf(
			s,
			"Expected non empty string, but got one with %d whitespace characters",
			len(s),
		)
	}

	return nil
}

// StringDigits validator for digits-only strings
type StringDigits string

// Validate performs validation
func (s StringDigits) Validate() error {
	if len(s) == 0 {
		return nil
	}

	for i, v := range []byte(s) {
		if v < 48 || v > 57 {
			return errorf(s, "Non number (char code: %d) found at position %d", v, i)
		}
	}

	return nil
}

// StringLatin is validation for latin characters-only strings
type StringLatin string

// Validate performs validation
func (s StringLatin) Validate() error {
	if len(s) == 0 {
		return nil
	}

	if !rxAlphaLatin.MatchString(string(s)) {
		return errorf(s, "Expected latin characters-only string")
	}

	return nil
}

// StringAlpha is validation for characters-only strings
// UTF-8 compatible, supports any language
type StringAlpha string

// Validate performs validation
func (s StringAlpha) Validate() error {
	if len(s) == 0 {
		return nil
	}

	if !rxAlpha.MatchString(string(s)) {
		return errorf(s, "Expected characters-only string")
	}

	return nil
}

// StringLuhn is validator, that used Luhn algorithm
// Primary usecase - credit card validation
type StringLuhn string

// Validate performs validation
func (s StringLuhn) Validate() error {
	if len(s) == 0 {
		return nil
	}
	if err := StringDigits(s).Validate(); err != nil {
		return err
	}

	var iv int
	bts := []byte(s)
	j := 0
	sum := 0
	for i := len(bts) - 1; i >= 0; i-- {
		iv = int(bts[i] - 48)
		if j%2 == 1 {
			iv *= 2
			if iv > 9 {
				iv = 1 + (iv % 10)
			}
		}

		sum += iv

		j++
	}

	if sum%10 != 0 {
		return errorf(s, "Lunh checksum is invalid - %d", sum)
	}

	return nil
}
