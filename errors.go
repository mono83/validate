package validate

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

// General validation error
type validationError struct {
	k reflect.Kind
	m string
}

func (v validationError) Error() string {
	return "[Validation failed on " + v.k.String() + "] " + v.m
}

// errorf builds and returns validation error
func errorf(src Interface, format string, a ...interface{}) validationError {
	return validationError{
		k: reflect.TypeOf(src).Kind(),
		m: fmt.Sprintf(format, a...),
	}
}

// ErrorsList represents list of (validation) errors
// Do not use it on other ways
type ErrorsList []error

func (e ErrorsList) Error() string {
	buf := bytes.NewBufferString("Multiple validation errors (")
	buf.WriteString(strconv.Itoa(len(e)))
	buf.WriteRune(')')

	for i, err := range e {
		if i > 0 {
			buf.WriteRune('\n')
		}
		buf.WriteString(err.Error())
	}

	return buf.String()
}
