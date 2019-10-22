package validate

import (
	"errors"
	"testing"
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
