.PHONY: help
help:
	@echo ""
	@echo "Commands:"
	@echo "    make test            # test function"
	@echo "    make lint            # lint function"
	@echo ""

.PHONY: test
test:
	vgo test ./...

.PHONY: lint
lint:
	vgo vet
	gometalinter \
		--enable-all \
		--deadline=300s \
		--line-length=105 \
		./...



