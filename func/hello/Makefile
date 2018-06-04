FUNC=https://my-awesome-api.example.com/

.PHONY: help
help:
	@echo "Func:   $(FUNC)"
	@echo ""
	@echo "Commands:"
	@echo "    make test            # test function"
	@echo "    make check           # checks function: $(FUNC)"
	@echo ""

.PHONY: test
test:
	go test

.PHONY: check
check:
	curl -XPOST -d "world" "$(FUNC)"
	curl -XPOST -d "Anders" "$(FUNC)"


