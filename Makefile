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

.PHONY: build-cross
build-cross: LDFLAGS += -extldflags "-static"
build-cross:
	CGO_ENABLED=0 gox -parallel=3 -output="_dist/{{.OS}}-{{.Arch}}/{{.Dir}}" -osarch='$(TARGETS)' $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' github.com/fishworks/gofish/cmd/$(NAME)

.PHONY: dist
dist:
	( \
		cd _dist && \
		$(DIST_DIRS) cp ../LICENSE {} \; && \
		$(DIST_DIRS) cp ../README.md {} \; && \
		$(DIST_DIRS) tar -zcf $(NAME)-${VERSION}-{}.tar.gz {} \; && \
		$(DIST_DIRS) zip -r $(NAME)-${VERSION}-{}.zip {} \; \
	)

.PHONY: checksum
checksum:
	for f in _dist/*.{gz,zip} ; do \
		shasum -a 256 "$${f}"  | awk '{print $$1}' > "$${f}.sha256" ; \
	done

.PHONY: test
test: TESTFLAGS += -race -v
test: test-lint test-unit

.PHONY: test-lint
test-lint:
	scripts/lint.sh

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
