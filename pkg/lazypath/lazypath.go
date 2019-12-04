package lazypath

import (
	"os"
	"path/filepath"
)

type LazyPath struct {
	EnvironmentVariable string
	DefaultFn           func() string
}

func (l LazyPath) Path(elem ...string) string {
	base := os.Getenv(l.EnvironmentVariable)
	if base == "" {
		base = l.DefaultFn()
	}
	return filepath.Join(base, filepath.Join(elem...))
}
