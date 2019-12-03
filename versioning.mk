MUTABLE_VERSION := canary

GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)

BINARY_VERSION ?= ${GIT_TAG}

# Only set Version if building a tag or VERSION is set
ifneq ($(BINARY_VERSION),)
	LDFLAGS += -X github.com/fishworks/gofish/version.Version=${BINARY_VERSION}
	VERSION = ${BINARY_VERSION}
else
	LDFLAGS += -X github.com/fishworks/gofish/version.BuildMetadata=git.${GIT_COMMIT}
	VERSION = ${MUTABLE_VERSION}
endif

info:
	 @echo "Version:           ${VERSION}"
	 @echo "Git Tag:           ${GIT_TAG}"
	 @echo "Git Commit:        ${GIT_COMMIT}"
