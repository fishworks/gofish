package osutil

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestExists(t *testing.T) {
	file, err := ioutil.TempFile("", "osutil")
	if err != nil {
		t.Fatal(err)
	}
	name := file.Name()
	// on Windows, we need to close all open handles to a file before we remove it.
	file.Close()

	exists, err := Exists(name)
	if err != nil {
		t.Errorf("expected no error when calling Exists() on a file that exists, got %v", err)
	}
	if !exists {
		t.Error("expected tempfile to exist")
	}
	os.Remove(name)
	stillExists, err := Exists(name)
	if err != nil {
		t.Errorf("expected no error when calling Exists() on a file that does not exist, got %v", err)
	}
	if stillExists {
		t.Error("expected tempfile to NOT exist after removing it")
	}
}
