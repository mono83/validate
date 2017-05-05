package validate

// IntPositive is validation container for positive integers
type IntPositive int

func (i IntPositive) Validate() error {
	if i < 1 {
		return errorf(i, "Expected positive value, but got %d", int(i))
	}

	return nil
}

// IntOdd is validation container for odd integers
type IntOdd int

func (i IntOdd) Validate() error {
	if i % 2 == 0 {
		return errorf(i, "Expected odd value, but got even %d", int(i))
	}

	return nil
}

// IntEven is validation container for even integers
type IntEven int

func (i IntEven) Validate() error {
	if i % 2 != 0 {
		return errorf(i, "Expected even value, but got odd %d", int(i))
	}

	return nil
}