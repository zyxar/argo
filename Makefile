GO          = go
PRODUCT     = argo
GOARCH     := amd64
# VERSION    := $(shell git describe --all --always --dirty --long)
# BUILD_TIME := $(shell date +%FT%T%z)
# GOVERSION  := $(shell go version | cut -d ' ' -f 3)
# LDFLAGS     = -ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.GoVersion=${GOVERSION}"

all: $(shell $(GO) env GOOS)

build:
	env GO111MODULE=$(GO111MODULE) GOOS=$(GOOS) GOARCH=$(GOARCH) GOFLAGS=$(GOFLAGS) $(GO) build $(GCFLAGS) -v -o $(PRODUCT)$(EXT) .

linux: export GOOS=linux
linux: EXT=.elf
linux: build

darwin: export GOOS=darwin
darwin: EXT=.mach
darwin: build

js: export GOOS=js
js: export GOARCH=wasm
js: EXT=.wasm
js: build

.PHONY: clean
clean:
	@rm -f $(PRODUCT) $(PRODUCT).*
