SHELL := /bin/bash
SUBPKGS = cli file json memory

.PHONY: deps test build

deps:
	@echo "Checking and updating dependencies"
	@if ! [ -x "$(GOBIN)/dep" ]; then \
		go get -u github.com/go/dep; \
	fi
	dep ensure -update

test:
	@echo "Running package tests"
	go test ./...

coverage: coverage-clean $(SUBPKGS) coverage-cinder
	@echo "Combining coverage files"
	@echo "mode: set" > cover.out
	@for file in $(shell find "./_cover" -type f -name "*out"); do \
		cat $$file | sed 's|mode: set||' | tee -a cover.out; \
	done

$(SUBPKGS):
	@echo "Building coverage profile for $@"
	@go test -coverprofile "./_cover/$@.out" "./handlers/$@" &>/dev/null

coverage-cinder:
	@echo "Building coverage profile for cinder"
	@go test -coverprofile "./_cover/cinder.out" &>/dev/null

coverage-clean:
	@echo "Cleaning old coverage profiles"
	@rm -rf "$$PWD/_cover"
	@mkdir -p "$$PWD/_cover"
