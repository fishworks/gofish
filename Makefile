BIN_NAME      ?= gofish

# go option
GO        ?= go
TAGS      :=
LDFLAGS   :=
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin
VENDORDIR := $(CURDIR)/vendor

# Required for globs to work correctly
SHELL=/bin/bash

.PHONY: all
all: $(BINDIR)/$(BIN_NAME)

$(BINDIR)/$(BIN_NAME): libmruby.a
	$(GO) build -o $(BINDIR)/$(BIN_NAME) $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' ./cmd/$(BIN_NAME)

libmruby.a:
	go mod vendor
	cd vendor/github.com/SeekingMeaning/go-mruby && make libmruby.a
	cp vendor/github.com/SeekingMeaning/go-mruby/libmruby.a .

.PHONY: clean
clean:
	@rm -rf $(BINDIR)
	@rm -rf $(VENDORDIR)
	@rm -rf libmruby.a

include versioning.mk
