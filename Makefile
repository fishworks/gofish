NAME              ?= gofish
TARGETS           ?= darwin/amd64 linux/amd64 linux/386 linux/arm linux/arm64 linux/ppc64le windows/amd64
DIST_DIRS         = find * -type d -exec

# go option
GO        ?= go
TAGS      :=
TESTS     := .
TESTFLAGS :=
LDFLAGS   :=
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin

# Required for globs to work correctly
SHELL=/bin/bash

.PHONY: all
all: build

.PHONY: build
build:
	$(GO) build -o $(BINDIR)/$(NAME) $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' ./cmd/$(NAME)

.PHONY: test
test: TESTFLAGS += -race -v
test: test-lint test-unit

.PHONY: test-unit
test-unit:
	$(GO) test $(GOFLAGS) -cover -run $(TESTS) ./... $(TESTFLAGS)

.PHONY: protoc
protoc:
	$(MAKE) -C _proto/ all

.PHONY: clean
clean:
	@rm -rf $(BINDIR) ./_dist

include versioning.mk
