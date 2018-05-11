package main

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/fishworks/gofish"
)

func TestTankFill(t *testing.T) {
	gofish.HomePath = "/usr/local/Fish"
	expectedTank := tank{
		"GOFISH_HOME":        "/usr/local/Fish",
		"GOFISH_BARREL":      "/usr/local/Fish/Barrel",
		"GOFISH_RIGS":        "/usr/local/Fish/Rigs",
		"GOFISH_DEFAULT_RIG": "/usr/local/Fish/Rigs/github.com/fishworks/fish-food",
	}

	switch runtime.GOOS {
	case "darwin":
		expectedTank["GOFISH_CACHE"] = filepath.Join(gofish.UserHomePath, "Library", "Caches", "Fish")
	case "linux":
		expectedTank["GOFISH_CACHE"] = filepath.Join(gofish.UserHomePath, ".gofish")
	case "windows":
		expectedTank["GOFISH_CACHE"] = filepath.Join(gofish.UserHomePath, "AppData", "Local", "Fish")
	}

	tank := tank{}
	tank.fill()

	if !reflect.DeepEqual(tank, expectedTank) {
		t.Errorf("expected tanks to be equal; got '%v', wanted '%v'", tank, expectedTank)
	}
}
