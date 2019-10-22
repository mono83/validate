package validate

// All validates all provided values and returns error on problems
func All(values ...Interface) error {
	list := NewErrorsList()

	for _, v := range values {
		if err := v.Validate(); err != nil {
			list.Add(err)
		}
	}

	return list
}
