package validate

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestErrorsList(t *testing.T) {
	es := NewErrorsList()
	if es.Len() != 0 {
		t.Errorf("Expected empty slice but got %d", es.Len())
		return
	}
	if es.OrNil() != nil {
		t.Errorf("Expected nil")
		return
	}

	es.Add(errors.New("foo"))
	es.Add(errors.New("bar"))
	if l := es.Len(); l != 2 {
		t.Errorf("Expected 2 but got %d", l)
	}
	if l := es.OrNil().(ErrorsList).Len(); l != 2 {
		t.Errorf("Expected 2 but got %d", l)
	}
}

var errorsListAssertDate = []struct {
	Given   interface{}
	Success bool
}{
	{nil, false},
	{0, false},
	{1, true},
	{-1, true},
	{float32(0), false},
	{float32(1), true},
	{float32(-1), true},
	{float64(0), false},
	{float64(1), true},
	{float64(-1), true},
	{"", false},
	{" ", true},
	{"foo", true},
	{time.Duration(0), false},
	{time.Millisecond, true},
	{IntPositive(0), false},  // Validable interface
	{IntPositive(-1), false}, // Validable interface
	{IntPositive(1), true},   // Validable interface
}

func TestErrorsList_Assert(t *testing.T) {
	for _, data := range errorsListAssertDate {
		t.Run(fmt.Sprintf("%T - %v", data.Given, data.Given), func(t *testing.T) {
			es := NewErrorsList()
			result := es.Assert(data.Given, "any")
			if result != data.Success {
				if result {
					t.Error("Expected failed assertion")
				} else {
					t.Error("Expected success assertion")
				}
			}
		})
	}
}

func TestErrorsList_AssertFuncReference(t *testing.T) {
	var holder struct {
		Func func(int)
	}

	es := NewErrorsList()
	es.Assert(holder.Func, "Func")
	if es.OrNil() == nil {
		t.Error("Expected error")
	}

	holder.Func = func(int) {
	}

	es = NewErrorsList()
	es.Assert(holder.Func, "Func")
	if err := es.OrNil(); err != nil {
		t.Error("Expected no error but got " + err.Error())
	}
}
