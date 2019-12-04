package main

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/fishworks/gofish/pkg/home"
)

func TestTankFill(t *testing.T) {
	os.Setenv(home.DefaultHomeEnvVar, "/usr/local/Fish")
	expectedTank := tank{
		"GOFISH_HOME":        "/usr/local/Fish",
		"GOFISH_BARREL":      "/usr/local/Fish/Barrel",
		"GOFISH_RIGS":        "/usr/local/Fish/Rigs",
		"GOFISH_DEFAULT_RIG": "/usr/local/Fish/Rigs/github.com/fishworks/fish-food",
	}

	if runtime.GOOS == "windows" {
		os.Setenv(home.DefaultHomeEnvVar, "C:\\Fish")
		expectedTank["GOFISH_HOME"] = "C:\\Fish"
		expectedTank["GOFISH_BARREL"] = "C:\\Fish\\Barrel"
		expectedTank["GOFISH_RIGS"] = "C:\\Fish\\Rigs"
		expectedTank["GOFISH_DEFAULT_RIG"] = "C:\\Fish\\Rigs\\github.com\\fishworks\\fish-food"
	}

	switch runtime.GOOS {
	case "darwin":
		expectedTank["GOFISH_CACHE"] = filepath.Join(home.UserHome(), "Library", "Caches", "gofish")
	case "linux":
		expectedTank["GOFISH_CACHE"] = filepath.Join(home.UserHome(), ".gofish")
	case "windows":
		expectedTank["GOFISH_CACHE"] = filepath.Join(home.UserHome(), "AppData", "Local", "gofish")
	}

	tank := tank{}
	tank.fill()

	if !reflect.DeepEqual(tank, expectedTank) {
		t.Errorf("expected tanks to be equal; got '%v', wanted '%v'", tank, expectedTank)
	}
}
