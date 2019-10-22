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

// NewErrorsList builds new errors list collection
func NewErrorsList() ErrorsList {
	return ErrorsList([]error{})
}

// ErrorsList represents list of (validation) errors
// Do not use it on other ways
type ErrorsList []error

func (e ErrorsList) Error() string {
	l := e.Len()
	if l == 0 {
		return "empty errors list - possible validate.ErrorsList misusage"
	} else if l == 1 {
		return e[0].Error()
	}

	buf := bytes.NewBufferString("Multiple validation errors (")
	buf.WriteString(strconv.Itoa(len(e)))
	buf.WriteRune(')')

	for _, err := range e {
		buf.WriteRune('\n')
		buf.WriteString(err.Error())
	}

	return buf.String()
}

// Add adds new error to errors list
func (e *ErrorsList) Add(err error) {
	if err != nil {
		*e = append(*e, err)
	}
}

// Addf creates new error and adds it to list
func (e *ErrorsList) Addf(format string, a ...interface{}) {
	e.Add(fmt.Errorf(format, a...))
}

// Len returns amount of items inside errors list
func (e ErrorsList) Len() int {
	return len(e)
}

// OrNil analyzes error list contents and returns nil if
// list contains no data. Otherwise, whole list is returned
func (e ErrorsList) OrNil() error {
	l := e.Len()
	if l == 0 {
		return nil
	} else if l == 1 {
		return e[0]
	}

	return e
}
