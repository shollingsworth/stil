.DEFAULT_GOAL := help
GOARCH := amd64
PROJECT := github.com/shollingsworth
PREFIX := stil
BUILD_DIR := build
# Run go tool dist list to see all os & arch combinations that are supported.
GOOS := linux  # put other OSes here as needed
ARCH := amd64 # what architectures to build for
ENTRYPOINT := main.go

help:
	@echo "Available commands:"
	@echo
	@cat Makefile | grep '^\w.*:$$' | cut -d ':' -f 1 | grep -v '^help$$'

fmt:
	@go fmt ./...


build: setup clean
	go mod download
	mkdir -p $(BUILD_DIR)/
	for goos in $(GOOS); do \
		if [ "$${goos}" = "windows" ]; then ext=".exe"; else ext="" ; fi; \
		for arch in $(ARCH); do \
			GOOS="$${goos}" GOARCH="$${arch}" go build \
				-trimpath \
				-o "build/$(PREFIX)-$${goos}-$${arch}$${ext}" \
				$(ENTRYPOINT) ; \
			echo $${goos}-$${arch}; \
		done \
	done
	ls -lh $(BUILD_DIR)

setup:
	go mod init $(PROJECT)/$(PREFIX) || true
	go mod tidy

install: build
	cp $(BUILD_DIR)/$(PREFIX)-* ~/sbin/stil
	stil completion zsh > ~/.oh-my-zsh/completions/_stil 


clean:
	rm -fv $(BUILD_DIR)/*
