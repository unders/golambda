.PHONY: help
help:
	@echo ""
	@echo "Commands:"
	@echo "    make install           # Install dependencies"
	@echo "    make log               # Show git log"
	@echo ""

.PHONY: install
install:
	# go get -u golang.org/x/vgo
	# go get github.com/awslabs/goformation
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: log
log:
	@git log --graph --oneline --decorate



