-include .env

# Default goal
.PHONY: help
.DEFAULT_GOAL := help

.PHONY: _check-root help clean run-cli run-webserver production.pack-cli \
	production.pack-webserver production.pack-target \
	production.pack-all production.pack-matrix

# Default OS and architecture
OS ?= $(shell go env GOOS)
ARCH ?= $(shell go env GOARCH)

# Go source config
MAIN_GO_FILE := main.go
OUTPUT_DIR := dist

# Package entry points
MAIN_PACKAGE_PATH_WEBSERVER := cmd/webserver
MAIN_PACKAGE_PATH_CLI := cmd/cli

# Binary output names
BINARY_NAME_WEBSERVER := http-server
BINARY_NAME_CLI := cli

# Target platforms (cross-compilation matrix)
OS_LIST := linux darwin windows
ARCH_LIST := amd64 arm64

# Matrix targets
MATRIX_BINARY_NAMES := $(foreach os,$(OS_LIST), \
  $(foreach arch,$(ARCH_LIST), \
	$(addprefix $(BINARY_NAME_WEBSERVER)_,$(os)_$(arch)$(if $(filter $(os),windows),.exe)) \
	$(addprefix $(BINARY_NAME_CLI)_,$(os)_$(arch)$(if $(filter $(os),windows),.exe)) \
  ) \
)
MATRIX_BINARY_PATHS := $(foreach binary,$(MATRIX_BINARY_NAMES),$(OUTPUT_DIR)/$(binary))

# Detect extension
EXT := $(if $(filter $(OS),windows),.exe)


# go_build: Macro to build a binary
define go_build
	GOOS=$(1) GOARCH=$(2) go build -ldflags='-s' -o=$(3) $(4)/$(MAIN_GO_FILE)
endef


# compress_upx: Macro to compress the binary using UPX
define compress_upx
	echo "[*] Compressing $(1) using UPX for target OS: $(2)"
	if [ "$(2)" = "darwin" ]; then \
		upx --best --lzma --force-macos $(1) || (echo "[*] \033[31m\033[1mUPX failed, skipping...\033[0m"; true); \
	else \
		upx --best --lzma $(1) || (echo "[*] \033[31m\033[1mUPX failed, skipping...\033[0m"; true); \
	fi
endef


# build_and_pack: Macro to build and pack a binary
define build_and_pack
	mkdir -p $(dir $1)
	binary_name=$$(basename $1); \
	noext_name=$${binary_name%%.*}; \
	os=$$(echo $$noext_name | cut -d'_' -f2); \
	arch=$$(echo $$noext_name | cut -d'_' -f3); \
	if echo $$noext_name | grep -q '^$(BINARY_NAME_WEBSERVER)_'; then \
		package_path=$(MAIN_PACKAGE_PATH_WEBSERVER); \
	else \
		package_path=$(MAIN_PACKAGE_PATH_CLI); \
	fi; \
	echo "[+] - [$$os/$$arch] - Packing $$binary_name ($$package_path/$(MAIN_GO_FILE))..."; \
	GOOS=$$os GOARCH=$$arch go build -ldflags='-s' -o=$1 $$package_path/$(MAIN_GO_FILE); \
	echo "[*] Compressing $$binary_name using UPX for target OS: $$os"; \
	if [ "$$os" = "darwin" ]; then \
		upx --best --lzma --force-macos $1 || (echo "[*] \033[31m\033[1mUPX failed, skipping...\033[0m"; true); \
	else \
		upx --best --lzma $1 || (echo "[*] \033[31m\033[1mUPX failed, skipping...\033[0m"; true); \
	fi
endef


# Pattern rule for matrix builds
$(OUTPUT_DIR)/%:
	@$(call build_and_pack,$@)


# Recipe for checking the project root
_check-root:
	@test -f go.mod || (echo "Please run this Makefile from the project root (go.mod missing)"; exit 1)


# Recipe for showing available targets
help: ## Show available make targets
	@grep -E '^[a-zA-Z0-9_/.-]+:.*?## .+$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2}'


# Recipe for cleaning the output directory
clean: ## Clean up the output directory
	@echo "[*] Cleaning up $(OUTPUT_DIR)"
	@rm -rf $(OUTPUT_DIR)


# Recipe for running the CLI
run-cli: _check-root ## Run the CLI
	@echo "[*] Running $(BINARY_NAME_CLI) for $(OS)/$(ARCH)"
	@GOOS=$(OS) GOARCH=$(ARCH) go run ./$(MAIN_PACKAGE_PATH_CLI)


# Recipe for running the webserver
run-webserver: _check-root ## Run the webserver
	@echo "[*] Running $(BINARY_NAME_WEBSERVER) for $(OS)/$(ARCH)"
	@GOOS=$(OS) GOARCH=$(ARCH) go run ./$(MAIN_PACKAGE_PATH_WEBSERVER)


# Recipe for building and packing the webserver
.DELETE_ON_ERROR:
production.pack-webserver: ## Build and pack the webserver binary
	@echo "[*] Building $(BINARY_NAME_WEBSERVER) for ${OS}/${ARCH}"
	$(call go_build,$(OS),$(ARCH),$(OUTPUT_DIR)/$(BINARY_NAME_WEBSERVER)$(EXT),$(MAIN_PACKAGE_PATH_WEBSERVER))
	@$(call compress_upx,$(OUTPUT_DIR)/$(BINARY_NAME_WEBSERVER)$(EXT),$(OS))

# Recipe for building and packing the CLI
.DELETE_ON_ERROR:
production.pack-cli: ## Build and pack the CLI binary
	@echo "[*] Building $(BINARY_NAME_CLI) for ${OS}/${ARCH}"
	$(call go_build,$(OS),$(ARCH),$(OUTPUT_DIR)/$(BINARY_NAME_CLI)$(EXT),$(MAIN_PACKAGE_PATH_CLI))
	@$(call compress_upx,$(OUTPUT_DIR)/$(BINARY_NAME_CLI)$(EXT),$(OS))

# Recipe for building and packing the CLI and webserver for a specific OS/ARCH
.DELETE_ON_ERROR:
production.pack-target: ## Build and pack the CLI and webserver binaries for a specific OS/ARCH
	@$(MAKE) $(OUTPUT_DIR)/$(BINARY_NAME_CLI)_$(OS)_$(ARCH)$(if $(filter $(OS),windows),.exe)
	@$(MAKE) $(OUTPUT_DIR)/$(BINARY_NAME_WEBSERVER)_$(OS)_$(ARCH)$(if $(filter $(OS),windows),.exe)


# Recipe for building and packing the CLI and webserver
.DELETE_ON_ERROR:
production.pack-all: production.pack-webserver production.pack-cli ## Build and pack CLI and webserver binaries for local OS/ARCH


# Recipe for cross-compiling and packing the CLI and webserver
.DELETE_ON_ERROR:
production.pack-matrix: clean $(MATRIX_BINARY_PATHS) ## Cross-compile and pack CLI and webserver binaries for all OS/ARCH combinations
