package version

import (
	"fmt"
	"strings"
)

// Version is the current version of GoFish.
var Version = ""

// BuildMetadata is the extra build time data.
var BuildMetadata = ""

// String represents the version information as a well-formatted string.
func String() string {
	ver := "dev"
	if strings.Compare(Version, "") != 0 {
		ver = fmt.Sprintf("%s", Version)
	}
	if strings.Compare(BuildMetadata, "") != 0 {
		ver = fmt.Sprintf("%s+%s", ver, BuildMetadata)
	}
	return ver
}
