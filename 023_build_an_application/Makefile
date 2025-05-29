.DEFAULT_GOAL := help

# ==================================================================================== #
# PACKAGE CONFIGURATION
# ==================================================================================== #

# Sets the path to the main package, binary name, and output directory
MAIN_PACKAGE_PATH := cmd/webserver
MAIN_GO_FILE := main.go
BINARY_NAME := http-server
OUTPUT_DIR := dist

# ==================================================================================== #
# SINGLE BINARY CONFIGURATION
# ==================================================================================== #

OS := linux
ARCH := amd64

# ==================================================================================== #
# CROSS-COMPILE CONFIGURATION (/matrix recipes)
# ==================================================================================== #

# Set the OS and ARCH options based for cross-compilation
OS_LIST := linux darwin windows
ARCH_LIST := amd64 arm64

# ===================================================================================== #

# Create a list of all possible OS/ARCH combinations
MATRIX_BINARY_NAMES := $(foreach os,$(OS_LIST),$(foreach arch,$(ARCH_LIST),$(BINARY_NAME)_$(os)_$(arch)))

# Create the full paths for the binaries
MATRIX_BINARY_PATHS := $(foreach binary,$(MATRIX_BINARY_NAMES),$(OUTPUT_DIR)/$(binary))

# Change the binary name depending on the OS
ifeq "$(OS)" "windows"
	BINARY_NAME := ${BINARY_NAME}_${OS}_${ARCH}.exe
endif
ifeq "$(OS)" "darwin"
	BINARY_NAME := ${BINARY_NAME}_${OS}_${ARCH}
endif
ifeq "$(OS)" "linux"
	BINARY_NAME := ${BINARY_NAME}_${OS}_${ARCH}
endif

# ==================================================================================== #
# ENTRYPOINT
# ==================================================================================== #

## help: print this help message (default target)
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
	
# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tidy: [development] tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	gofumpt -l -w .

## run: [development] run the application
.PHONY: run
run:
	GOOS=${OS} GOARCH=${ARCH} go run ${MAIN_PACKAGE_PATH}

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: [quality-control] run quality control checks
.PHONY: audit
audit:
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofumpt -l .)" 
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: [quality-control] run all tests
.PHONY: test
test:
	go test -v -race -buildvcs=true ./...

## test/cover: [quality-control] run all tests and display coverage
.PHONY: test/cover
test/cover: coverage
	go test -v -race -buildvcs=true -coverprofile=${OUTPUT_DIR}/coverage/coverage.out ./...
	go tool cover -html=${OUTPUT_DIR}/coverage/coverage.out -o ${OUTPUT_DIR}/coverage/coverage.html

## upgradeable: [quality-control] list direct dependencies that have upgrades available
.PHONY: upgradeable
upgradeable:
	@go list -u -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -m all

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

# no-dirty: [helper] ensure the working directory is clean before running commands
.PHONY: no-dirty
no-dirty:
	@if [ -n "$(shell git status --porcelain)" ]; then \
		echo "Error: Working directory is not clean. Please commit or stash your changes."; \
		exit 1; \
	fi

# coverage: [helper] create a directory for coverage reports
.PHONY: coverage
coverage:
	mkdir -p $@

# clean-cache: [helper] remove all cached files
.PHONY: clean-cache
clean-cache:
	@echo "Cleaning up cached files..."
	@go clean -cache -testcache -modcache

## build: [helper] build the application
.PHONY: build
build: clean
	@echo "> [${OS}/${ARCH}] - Building ${BINARY_NAME} (${MAIN_PACKAGE_PATH}/${MAIN_GO_FILE})..."
	GOOS=${OS} GOARCH=${ARCH} go build -o=${OUTPUT_DIR}/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}/${MAIN_GO_FILE}

## clean: [helper] remove all build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up build artifacts..."
	@rm -rf ${OUTPUT_DIR}
	@rm -rf coverage
	@rm -rf *.db*

# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

## build/matrix: [operations] build the application for all OS/ARCH combinations
.PHONY: build/matrix
build/matrix: $(MATRIX_BINARY_PATHS)
$(OUTPUT_DIR)/%:
	$(MAKE) build OUTPUT_DIR=$(dir $@) BINARY_NAME=$(notdir $@) OS=$(word 2,$(subst _, ,$(notdir $@))) ARCH=$(word 3,$(subst _, ,$(notdir $@)))

## production/pack: [operations] pack (compress) the application for production
.PHONY: production/pack 
production/pack: audit no-dirty clean-cache
	GOOS=${OS} GOARCH=${ARCH} go build -ldflags='-s' -o=${OUTPUT_DIR}/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}/${MAIN_GO_FILE}

	@if [ "${OS}" = "darwin" ]; then \
		upx --best --lzma --force-macos ${OUTPUT_DIR}/bin/${BINARY_NAME} || (echo "> \033[31m\033[1mUPX failed for ${OS}/${ARCH}, skipping...\033[0m"; true); \
	else \
		upx --best --lzma ${OUTPUT_DIR}/bin/${BINARY_NAME} || (echo "> \033[31m\033[1mUPX failed for ${OS}/${ARCH}, skipping...\033[0m"; true); \
	fi

## production/pack/matrix: [operations] pack (compress) the application for production for all OS/ARCH combinations
.PHONY: production/pack/matrix
production/pack/matrix: $(MATRIX_BINARY_PATHS)
$(OUTPUT_DIR)/%:
	@$(MAKE) production/pack OUTPUT_DIR=$(dir $@) BINARY_NAME=$(notdir $@) OS=$(word 2,$(subst _, ,$(notdir $@))) ARCH=$(word 3,$(subst _, ,$(notdir $@)))
