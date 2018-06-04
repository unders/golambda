FUNC=https://my-awesome-api.example.com/

.PHONY: help
help:
	@echo "Func:   $(FUNC)"
	@echo ""
	@echo "Commands:"
	@echo "    make test            # test function"
	@echo "    make lint            # lint function"
	@echo "    make check           # checks function: $(FUNC)"
	@echo ""


.PHONY: build
build:
	GOOS=linux vgo build -o main
	zip deployment.zip main

.PHONY: test
test:
	vgo test

.PHONY: lint
lint:
	vgo vet
	gometalinter \
		--enable-all \
		--deadline=300s \
		--line-length=102 \
		./...

.PHONY: check
check:
	curl -XPOST -d "world" "$(FUNC)"
	curl -XPOST -d "Anders" "$(FUNC)"


