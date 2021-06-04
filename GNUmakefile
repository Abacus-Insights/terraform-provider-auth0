PKG_NAME = auth0
PKGS ?= $$(go list ./...)
FILES ?= $$(find . -name '*.go' | grep -v vendor)
TESTS ?= ".*"
COVERS ?= "c.out"
SUPPORTED_ARCH = linux/amd64 darwin/amd64 darwin/arm64
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(shell dirname $(mkfile_path))
PLUGIN_DIR := ~/.terraform.d/plugins
	
default: build

build: fmtcheck modedit
	@go install

buildall: fmtcheck modedit
	@GOOS=linux GOARCH=amd64 go build -o $(current_dir)/terraform-provider-$(PKG_NAME)-linux
	@GOOS=darwin GOARCH=amd64 go build -o $(current_dir)/terraform-provider-$(PKG_NAME)-darwin

install: build
	@mkdir -p $(PLUGIN_DIR)
	@cp $(GOPATH)/bin/terraform-provider-auth0 $(PLUGIN_DIR)

installall: buildall
	@mkdir -p $(PLUGIN_DIR)/{darwin,linux}_amd64
	@cp $(current_dir)/terraform-provider-$(PKG_NAME)-linux $(PLUGIN_DIR)/linux_amd64/terraform-provider-$(PKG_NAME)
	@cp $(current_dir)/terraform-provider-$(PKG_NAME)-darwin $(PLUGIN_DIR)/darwin_amd64/terraform-provider-$(PKG_NAME)


sweep:
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	@go test ./auth0 -v -sweep="phony" $(SWEEPARGS)

test: fmtcheck
	@go test -i $(PKGS) || exit 1
	@echo $(PKGS) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4 -run ^$(TESTS)$

testacc: fmtcheck
	@TF_ACC=1 go test $(PKGS) -v $(TESTARGS) -timeout 120m -coverprofile=$(COVERS) -run ^$(TESTS)$

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	@gofmt -w $(FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

modedit:
	@go mod edit -replace="gopkg.in/auth0.v5=$(current_dir)/../auth0"

docgen:
	go run scripts/gendocs.go -resource auth0_<resource>

.PHONY: build test testacc vet fmt fmtcheck errcheck docgen
