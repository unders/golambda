##
## BUILD
##
VERSION=1
MAIN_VERSION="main.Version=v$(VERSION)"
BUILDSTAMP="main.Buildstamp=$(shell date -u '+%Y-%m-%dT%I:%M%p')"
GITHASH="main.Githash=$(shell git rev-parse HEAD)"
LDFLAGS=-ldflags "-X $(MAIN_VERSION) -X $(BUILDSTAMP) -X $(GITHASH)"

##
## DEPLOY
##
ACCOUNT_ID=613705717678
ROLE=Lambda

##
## ENDPOINT
##
FUNC=https://my-awesome-api.example.com/

.PHONY: help
help:
	@echo "Func:   $(FUNC)"
	@echo ""
	@echo "Commands:"
	@echo "    make build            # build function"
	@echo "    make create           # create lambda function"
	@echo "    make check            # checks function: $(FUNC)"
	@echo "    make clean            # remove build artifacts"
	@echo ""

.PHONY: build
build: clean test lint
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 vgo build $(LDFLAGS)
	zip deployment.zip hello

.PHONY: create
create: build
	aws lambda create-function \
		--region us-west-1 \
		--function-name HelloFunction \
		--zip-file fileb://./deployment.zip \
		--runtime go1.x \
		--tracing-config Mode=Active \
		--role arn:aws:iam::$(ACCOUNT_ID):role/$(ROLE) \
		--handler hello \
		--profile underscli
	make clean

.PHONY: check
check:
	curl -XPOST -d "world" "$(FUNC)"
	curl -XPOST -d "Anders" "$(FUNC)"

.PHONY: clean
clean:
	rm hello
	rm deployment.zip

##
##  Private
##

.PHONY: test
test:
	cd ../../lambda/hello && make test

.PHONY: lint
lint:
	cd ../../lambda/hello && make lint

