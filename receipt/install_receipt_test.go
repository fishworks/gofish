package receipt

import (
	"bytes"
	"testing"
)

func TestNewFromReader(t *testing.T) {
	buf := bytes.NewBufferString("{}")
	r, err := NewFromReader(buf)
	if err != nil {
		t.Error(err)
	}
	expected := InstallReceipt{}
	if *r != expected {
		t.Errorf("expected to load an empty install receipt")
	}

	buf = bytes.NewBufferString(`{"name": "foo"}`)
	r, err = NewFromReader(buf)
	if err != nil {
		t.Error(err)
	}
	expected.Name = "foo"
	if *r != expected {
		t.Errorf("expected '%v', got '%v'", expected, *r)
	}
}
