package validate

import "errors"

type replaceMessage struct {
	v Interface
	m string
}

func (a replaceMessage) Validate() error {
	err := a.v.Validate()
	if err == nil {
		return nil
	}

	if e, ok := err.(validationError); ok {
		return validationError{
			k: e.k,
			m: a.m,
		}
	}

	return errors.New(a.m)
}

// WithMessage wraps validator with custom message provider
func WithMessage(v Interface, m string) Interface {
	return replaceMessage{v: v, m: m}
}
