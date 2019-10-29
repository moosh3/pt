.PHONY: all clean

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

REGISTRY=r.kubernetes.lol
BINARY=pt
VERSION=0.1.0
BUILD=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

all: clean
	@echo "Building binary version: $(VERSION), build: $(BUILD)"
	@go build -o $(BINARY) $(LDFLAGS)

docker:
	@echo 'Building docker image version: $(VERSION), build: $(BUILD)'
	@docker build -t "$(REGISTRY)/$(BINARY):$(VERSION)" \
		--build-arg build=$(BUILD) --build-arg version=$(VERSION) \
		-f Dockerfile .

clean:
	@echo 'Removing present binaries'
	@-rm $(BINARY) >/dev/null
