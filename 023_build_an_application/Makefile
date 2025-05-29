.DEFAULT_GOAL := help

# ==================================================================================== #
# PACKAGE CONFIGURATION
# ==================================================================================== #

# Sets the path to the main package and the main Go file
MAIN_PACKAGE_PATH := cmd/webserver
MAIN_GO_FILE := main.go

# Sets the binary name and output directory
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
MATRIX_BINARY_NAMES := $(foreach os,$(OS_LIST),$(foreach arch,$(ARCH_LIST),$(BINARY_NAME)_$(os)_$(arch)$(if $(filter $(os),windows),.exe)))

# Create the full paths for the binaries
MATRIX_BINARY_PATHS := $(foreach binary,$(MATRIX_BINARY_NAMES),$(OUTPUT_DIR)/$(binary))

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
	GOOS=${OS} GOARCH=${ARCH} go run ${MAIN_PACKAGE_PATH}/${MAIN_GO_FILE}

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
test/cover: coverage clean-test-cache
	@mkdir -p ${OUTPUT_DIR}/coverage
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
		echo "[+] Error: Working directory is not clean. Please commit or stash your changes."; \
		exit 1; \
	fi

# coverage: [helper] create a directory for coverage reports
.PHONY: coverage
coverage:
	@mkdir -p ${OUTPUT_DIR}/$@

# clean-test-cache: [helper] remove all test cache files
.PHONY: clean-test-cache
clean-test-cache:
	@echo "[+] Cleaning up test cache files..."
	@go clean -testcache

## clean: [helper] remove all build artifacts
.PHONY: clean
clean:
	@echo "[+] Cleaning up build artifacts..."
	@rm -rf ${OUTPUT_DIR}
	@rm -rf coverage
	@rm -rf *.db*

# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

## build: [operations] build the application
.PHONY: build
.DELETE_ON_ERROR:
build: no-dirty audit test
	@if [ "${OS}" = "windows" ]; then \
		BINARY_NAME="${BINARY_NAME}_${OS}_${ARCH}.exe"; \
	elif [ "${OS}" = "darwin" ]; then \
		BINARY_NAME="${BINARY_NAME}_${OS}_${ARCH}"; \
	else \
		BINARY_NAME="${BINARY_NAME}_${OS}_${ARCH}"; \
	fi

	@echo "[+] - [${OS}/${ARCH}] - Building ${BINARY_NAME} (${MAIN_PACKAGE_PATH}/${MAIN_GO_FILE})..."

	GOOS=${OS} GOARCH=${ARCH} go build -o=${OUTPUT_DIR}/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}/${MAIN_GO_FILE}

## production/pack: [operations] pack (compress) the application for production
.PHONY: production/pack 
.DELETE_ON_ERROR:
production/pack: no-dirty clean audit
	@echo "[+] - [${OS}/${ARCH}] - Packing ${BINARY_NAME} (${MAIN_PACKAGE_PATH}/${MAIN_GO_FILE})..."

	@GOOS=${OS} GOARCH=${ARCH} go build -ldflags='-s' -o=${OUTPUT_DIR}/bin/${BINARY_NAME}$(if $(filter $(OS),windows),.exe) ${MAIN_PACKAGE_PATH}/${MAIN_GO_FILE}

	@if [ "${OS}" = "darwin" ]; then \
		upx --best --lzma --force-macos ${OUTPUT_DIR}/bin/${BINARY_NAME}$(if $(filter $(OS),windows),.exe) || (echo "[*] \033[31m\033[1mUPX failed for ${OS}/${ARCH}, skipping...\033[0m"; true); \
	else \
		upx --best --lzma ${OUTPUT_DIR}/bin/${BINARY_NAME}$(if $(filter $(OS),windows),.exe) || (echo "[*] \033[31m\033[1mUPX failed for ${OS}/${ARCH}, skipping...\033[0m"; true); \
	fi

## production/pack/matrix: [operations] pack (compress) the application for production for all OS/ARCH combinations
.PHONY: production/pack/matrix
.DELETE_ON_ERROR:
production/pack/matrix: $(MATRIX_BINARY_PATHS) test
$(OUTPUT_DIR)/%:
	@echo "[+] - [$(word 2,$(subst _, ,$(notdir $@)))/$(word 3,$(subst _, ,$(basename $(notdir $@))))] - Packing $(BINARY_NAME) ($(MAIN_PACKAGE_PATH)/$(MAIN_GO_FILE))..."

	@$(MAKE) production/pack OUTPUT_DIR=$(dir $@) BINARY_NAME=$(notdir $@) OS=$(word 2,$(subst _, ,$(notdir $@))) ARCH=$(word 3,$(subst _, ,$(basename $(notdir $@))))
