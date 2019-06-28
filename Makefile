GO          = go
GOARCH      = amd64
GO111MODULE = on
GOFLAGS     =
SRC         =
PRODUCT     = argo

all: $(shell $(GO) env GOOS)

build:
	env GO111MODULE=$(GO111MODULE) GOOS=$(GOOS) GOARCH=$(GOARCH) GOFLAGS=$(GOFLAGS) $(GO) build ${LDFLAGS} $(GCFLAGS) -v -o bin/$(GOOS)_$(GOARCH)/$(PRODUCT) ./$(SRC)

install:
	env GO111MODULE=$(GO111MODULE) GOOS=$(GOOS) GOARCH=$(GOARCH) GOFLAGS=$(GOFLAGS) $(GO) install ./$(SRC)


linux: export GOOS=linux
linux: build

darwin: export GOOS=darwin
darwin: build

#build-ldflags: revision=$(shell git describe --all --always --dirty --long)
#build-ldflags: timestamp=$(shell date +%FT%T%z)
#build-ldflags: LDFLAGS=-ldflags "-X main.Revision=${revision} -X main.BuildTime=${timestamp}"

js: export GOOS=js
js: export GOARCH=wasm
js: build

.PHONY: clean gomod
clean:
	@rm -fr bin/*

gomod:
	env GO111MODULE=on $(GO) mod tidy
#	env GO111MODULE=on $(GO) mod vendor
