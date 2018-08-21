GO          = go
PRODUCT     = argo
GOARCH     := amd64
# VERSION    := $(shell git describe --all --always --dirty --long)
# BUILD_TIME := $(shell date +%FT%T%z)
# GOVERSION  := $(shell go version | cut -d ' ' -f 3)
# LDFLAGS     = -ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.GoVersion=${GOVERSION}"

all: $(shell $(GO) env GOOS)

build-%:
	$(eval $@_OS := $*)
	env GOOS=$($@_OS) GOARCH=$(GOARCH) $(GO) build ${LDFLAGS} -v -o $(PRODUCT)$(EXT) .


linux: EXT=.elf
linux: build-linux

darwin: EXT=.mach
darwin: build-darwin

.PHONY: clean
clean:
	@rm -f $(PRODUCT) $(PRODUCT).elf $(PRODUCT).mach
