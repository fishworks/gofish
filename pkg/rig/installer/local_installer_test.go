package installer

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/fishworks/gofish/pkg/home"
)

var _ Installer = new(LocalInstaller)

func TestLocalInstaller(t *testing.T) {
	dh, err := ioutil.TempDir("", "fish-home-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dh)
	os.Setenv(home.DefaultHomeEnvVar, dh)

	if err := os.MkdirAll(home.Rigs(), 0755); err != nil {
		t.Fatalf("Could not create %s: %s", home.Rigs(), err)
	}

	source := "testdata/fish-food"
	i, err := New(source, "", "")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := Install(i); err != nil {
		t.Error(err)
	}

	expectedPath := filepath.Join(home.String(), "Rigs", "fish-food")
	if i.Path() != expectedPath {
		t.Errorf("expected path '%s', got %q", expectedPath, i.Path())
	}
}
