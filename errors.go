package validate

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"time"
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

type multipleErrorsContainer interface {
	List() []error
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

// Assert performs value assertion
// It will generate an error and add it to errors list if:
// - nil
// - int - value is zero
// - float32, float64 - value is zero
// - string - empty string
// - validate.Interface - error inside Validate() func
//
// Returns true if assertion was successful, false on error
func (e *ErrorsList) Assert(value interface{}, field string) bool {
	if value == nil {
		e.Addf(`"%s" expected to be not nil`, field)
		return false
	}

	// Value implements validable interface
	if v, ok := value.(Interface); ok {
		err := v.Validate()
		if err != nil {
			e.Add(err)
			return false
		}
		return true
	}

	switch value.(type) {
	case int:
		if value.(int) == 0 {
			e.Addf(`"%s" expected to be not empty, but zero found`, field)
			return false
		}
	case float32:
		if value.(float32) == 0 {
			e.Addf(`"%s" expected to be not empty, but zero found`, field)
			return false
		}
	case float64:
		if value.(float64) == 0 {
			e.Addf(`"%s" expected to be not empty, but zero found`, field)
			return false
		}
	case string:
		if value.(string) == "" {
			e.Addf(`"%s" expected to be not empty, but empty string found`, field)
			return false
		}
	case time.Duration:
		if value.(time.Duration) == 0 {
			e.Addf(`"%s" expected to be not empty, but empty duration found`, field)
			return false
		}
	default:
		refValue := reflect.ValueOf(value)
		if refValue.Kind() == reflect.Func {
			if refValue.IsNil() {
				e.Addf(`"%s" expected to be not nil`, field)
				return false
			}
			// Non nil reference to function
			return true
		}

		e.Addf("unable to perform assertion on %T for %s", value, field)
		return false
	}

	return true
}

// Add adds new error to errors list
func (e *ErrorsList) Add(err error) {
	if err != nil {
		// Checking against container
		if container, ok := err.(multipleErrorsContainer); ok {
			for _, err := range container.List() {
				e.Add(err)
			}
		} else {
			*e = append(*e, err)
		}
	}
}

// List returns list of errors stored inside list
func (e ErrorsList) List() []error {
	return e
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
