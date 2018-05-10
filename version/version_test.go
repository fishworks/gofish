package version

import "testing"

func TestString(t *testing.T) {
	if String() != "dev" {
		t.Errorf("Expected .String() to return 'dev', got '%s'", String())
	}
	Version = "0.1.0"
	if String() != "0.1.0" {
		t.Errorf("Expected .String() to return '0.1.0', got '%s'", String())
	}
	BuildMetadata = "windows.amd64"
	if String() != "0.1.0+windows.amd64" {
		t.Errorf("Expected .String() to return '0.1.0+windows.amd64', got '%s'", String())
	}
	Version = ""
	if String() != "dev+windows.amd64" {
		t.Errorf("Expected .String() to return 'dev+windows.amd64', got '%s'", String())
	}
}
