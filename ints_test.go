package validate

import (
	"testing"
)

func TestIntPositive_Validate(t *testing.T) {
	dataSetNoError(
		t,
		func(v interface{}) error { return IntPositive(v.(int)).Validate() },
		[]interface{}{1, 1000},
	)
	dataSetError(
		t,
		func(v interface{}) error { return IntPositive(v.(int)).Validate() },
		[]interface{}{0, -1, -1000},
	)
}

func TestIntOdd_Validate(t *testing.T) {
	dataSetNoError(
		t,
		func(v interface{}) error { return IntOdd(v.(int)).Validate() },
		[]interface{}{1, 3, -23},
	)
	dataSetError(
		t,
		func(v interface{}) error { return IntOdd(v.(int)).Validate() },
		[]interface{}{0, 2, 1024},
	)
}

func TestIntEven_Validate(t *testing.T) {
	dataSetNoError(
		t,
		func(v interface{}) error { return IntEven(v.(int)).Validate() },
		[]interface{}{0, 2, -30},
	)
	dataSetError(
		t,
		func(v interface{}) error { return IntEven(v.(int)).Validate() },
		[]interface{}{1, -1, 101},
	)
}