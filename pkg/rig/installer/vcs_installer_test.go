package installer

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/fishworks/gofish"
	"github.com/fishworks/gofish/pkg/rig"
)

var _ Installer = new(VCSInstaller)

func TestVCSInstallerSuccess(t *testing.T) {
	dh, err := ioutil.TempDir("", "fish-home-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dh)

	home := gofish.Home(dh)
	if err := os.MkdirAll(home.Rigs(), 0755); err != nil {
		t.Fatalf("Could not create %s: %s", home.Rigs(), err)
	}

	source := "https://github.com/fishworks/fish-food"
	i, err := New(source, "", "", home)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	expectedPath := home.Path("Rigs", "github.com", "fishworks", "fish-food")
	if i.Path() != expectedPath {
		t.Errorf("expected path '%s', got %q", expectedPath, i.Path())
	}

	// ensure a VCSInstaller was returned
	vi, ok := i.(*VCSInstaller)
	if !ok {
		t.Error("expected a VCSInstaller")
	}

	expectedName := "github.com/fishworks/fish-food"
	if vi.Name != expectedName {
		t.Errorf("expected name '%s', got '%s'", expectedName, vi.Name)
	}

	vi.Name = "foo"

	expectedPath = home.Path("Rigs", "foo")
	if i.Path() != expectedPath {
		t.Errorf("expected path '%s', got %q", expectedPath, i.Path())
	}

	if err := Install(i); err != nil {
		t.Error(err)
	}
}

func TestVCSInstallerUpdate(t *testing.T) {
	dh, err := ioutil.TempDir("", "fish-home-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dh)

	home := gofish.Home(dh)
	if err := os.MkdirAll(home.Rigs(), 0755); err != nil {
		t.Fatalf("Could not create %s: %s", home.Rigs(), err)
	}

	source := "https://github.com/fishworks/fish-food"
	i, err := New(source, "", "", home)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	// ensure a VCSInstaller was returned
	_, ok := i.(*VCSInstaller)
	if !ok {
		t.Error("expected a VCSInstaller")
	}

	if err := Update(i); err == nil {
		t.Error("expected error for rig does not exist, got none")
	} else if err.Error() != "rig does not exist" {
		t.Errorf("expected error for rig does not exist, got (%v)", err)
	}

	// Install rig before update
	if err := Install(i); err != nil {
		t.Error(err)
	}

	// Update rig
	if err := Update(i); err != nil {
		t.Error(err)
	}

	// Test update failure
	os.Remove(filepath.Join(i.Path(), "LICENSE"))
	// Testing update for error
	if err := Update(i); err == nil {
		t.Error("expected error for rig modified, got none")
	} else if err != rig.ErrRepoDirty {
		t.Errorf("expected error for rig modified, got (%v)", err)
	}

}
