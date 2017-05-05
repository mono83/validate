package validate

// All validates all provided values and returns error on problems
func All(values ...Interface) error {
	errors := []error{}
	for _, v := range values {
		if err := v.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) == 0 {
		return nil
	} else if len(errors) == 1 {
		return errors[0]
	}

	return ErrorsList(errors)
}
