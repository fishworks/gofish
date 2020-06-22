BIN_NAME      ?= gofish

# go option
GO        ?= go
TAGS      :=
LDFLAGS   :=
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin

# Required for globs to work correctly
SHELL=/bin/bash

.PHONY: all
all: $(BINDIR)/$(BIN_NAME)

$(BINDIR)/$(BIN_NAME):
	$(GO) build -o $(BINDIR)/$(BIN_NAME) $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' ./cmd/$(BIN_NAME)

.PHONY: clean
clean:
	@rm -rf $(BINDIR)

include versioning.mk
