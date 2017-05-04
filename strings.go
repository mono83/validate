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

// StringAlpha is validation for latin characters-only strings
type StringLatin string

// Validate performs validation
func (s StringLatin) Validate() error {
	if len(s) == 0 {
		return errorf(s, "Non empty string expected for latin alpha check")
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
		return errorf(s, "Non empty string expected for alpha check")
	}

	if !rxAlpha.MatchString(string(s)) {
		return errorf(s, "Expected characters-only string")
	}

	return nil
}
