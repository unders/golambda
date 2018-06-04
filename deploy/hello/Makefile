VERSION=1
MAIN_VERSION="main.Version=v$(VERSION)"
BUILDSTAMP="main.Buildstamp=$(shell date -u '+%Y-%m-%dT%I:%M%p')"
GITHASH="main.Githash=$(shell git rev-parse HEAD)"
LDFLAGS=-ldflags "-X $(MAIN_VERSION) -X $(BUILDSTAMP) -X $(GITHASH)"

FUNC=https://my-awesome-api.example.com/

.PHONY: help
help:
	@echo "Func:   $(FUNC)"
	@echo ""
	@echo "Commands:"
	@echo "    make build            # build function"
	@echo "    make check           # checks function: $(FUNC)"
	@echo ""

.PHONY: build
build: # test lint
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 vgo build $(LDFLAGS)
	zip deployment.zip hello

.PHONY: check
check:
	curl -XPOST -d "world" "$(FUNC)"
	curl -XPOST -d "Anders" "$(FUNC)"

##
##  Private
##

.PHONY: test
test:
	cd ../../func/hello && make test

.PHONY: lint
lint:
	cd ../../func/hello && make lint

