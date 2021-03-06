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
API=https://9ipt7e6fc5.execute-api.eu-central-1.amazonaws.com/staging/HelloFunction

.PHONY: help
help:
	@echo "Func:   $(FUNC)"
	@echo ""
	@echo "Commands:"
	@echo "    make build            # build function"
	@echo "    make create           # create lambda function"
	@echo "    make check            # checks function: $(FUNC)"
	@echo "    make role/open        # open aws roles in web browser"
	@echo "    make clean            # remove build artifacts"
	@echo ""

.PHONY: build
build: clean test lint
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 vgo build $(LDFLAGS)
	zip deployment.zip hello

.PHONY: create
create: build
	aws lambda create-function \
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
	curl -XPOST -d "world" "$(API)"
	@echo ""
	curl -XPOST -d "Anders" "$(API)"
	@echo ""

.PHONY: role/open
role/open:
	@open https://console.aws.amazon.com/iam/home#/roles/Lambda?section=permissions

.PHONY: clean
clean:
	@if [ -f hello ]; then rm hello; fi;
	@if [ -f deployment.zip ]; then rm deployment.zip; fi;

##
##  Private
##

.PHONY: test
test:
	cd ../../lambda/hello && make test

.PHONY: lint
lint:
	cd ../../lambda/hello && make lint

