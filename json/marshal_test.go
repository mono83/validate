package json

import (
	"github.com/mono83/validate"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	var v validate.IntPositive
	if err := Unmarshal([]byte(`12345`), &v); err != nil {
		t.Fatal(err)
	}
	if err := Unmarshal([]byte(`-12345`), &v); err == nil {
		t.Fatal("Error expected due positive-only validation")
	}
}
