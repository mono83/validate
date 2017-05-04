package validate

import (
	"testing"
)

// dataSetNoError is helper function to test validators
func dataSetNoError(t *testing.T, f func(interface{}) error, data []interface{}) {
	for i, v := range data {
		t.Logf("[%d] Passing %+v", i, v)
		if err := f(v); err != nil {
			t.Errorf("[%d] Value %T %+v failed with error %s", i, v, v, err)
		}
	}
}

// dataSetError is helper function to test validators
func dataSetError(t *testing.T, f func(interface{}) error, data []interface{}) {
	for i, v := range data {
		t.Logf("[%d] Passing %+v", i, v)
		if err := f(v); err == nil {
			t.Errorf("[%d] Value %T %+v succeded, but failure expected", i, v, v)
		}
	}
}
