WORKPLACE := $(shell pwd)

build: clean
	@mkdir -p $(WORKPLACE)/build
	go build -o $(WORKPLACE)/build

install: build
	@mkdir -p /usr/local/bin
	install $(WORKPLACE)/build/socket_cat /usr/local/bin

.PHONY: clean dev

clean: clean-deps
	rm -rf $(WORKPLACE)/build

clean-deps:
	go clean -cache
