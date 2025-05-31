# Config
MAIN_GO_FILE := main.go
OUTPUT_DIR := dist

# Package paths
MAIN_PACKAGE_PATH_WEBSERVER := cmd/webserver
MAIN_PACKAGE_PATH_CLI := cmd/cli

# Binary names
BINARY_NAME_WEBSERVER := http-server
BINARY_NAME_CLI := cli

# Cross-compilation matrix
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

# Default goal
.DEFAULT_GOAL := help

# build-and-pack: Internal helper function
define build_and_pack
	@mkdir -p $(dir $1)
	@binary_name=$$(basename $1); \
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
	if [ "$$os" = "darwin" ]; then \
		upx --best --lzma --force-macos $1 || (echo "[*] \033[31m\033[1mUPX failed for $$os/$$arch, skipping...\033[0m"; true); \
	else \
		upx --best --lzma $1 || (echo "[*] \033[31m\033[1mUPX failed for $$os/$$arch, skipping...\033[0m"; true); \
	fi
endef

.PHONY: compress-upx
compress-upx:
	@if [ "$(OS)" = "darwin" ]; then \
		upx --best --lzma --force-macos $(FILE) || (echo "[*] \033[31m\033[1mUPX failed, skipping...\033[0m"; true); \
	else \
		upx --best --lzma $(FILE) || (echo "[*] \033[31m\033[1mUPX failed, skipping...\033[0m"; true); \
	fi

.PHONY: help
help: ## Show available make targets
	@grep -E '^[a-zA-Z0-9_/.-]+:.*?## .+$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: production/pack-webserver
production/pack-webserver: ## Build and pack the webserver binary
	@echo "[*] Building $(BINARY_NAME_WEBSERVER) for ${OS}/${ARCH}"
	@GOOS=$(OS) GOARCH=$(ARCH) go build -ldflags='-s' -o=$(OUTPUT_DIR)/$(BINARY_NAME_WEBSERVER)$(EXT) $(MAIN_PACKAGE_PATH_WEBSERVER)/$(MAIN_GO_FILE)
	@$(MAKE) compress-upx FILE=$(OUTPUT_DIR)/$(BINARY_NAME_WEBSERVER)$(EXT)

.PHONY: production/pack-cli
production/pack-cli: ## Build and pack the CLI binary
	@echo "[*] Building $(BINARY_NAME_CLI) for ${OS}/${ARCH}"
	@GOOS=$(OS) GOARCH=$(ARCH) go build -ldflags='-s' -o=$(OUTPUT_DIR)/$(BINARY_NAME_CLI)$(EXT) $(MAIN_PACKAGE_PATH_CLI)/$(MAIN_GO_FILE)
	@$(MAKE) compress-upx FILE=$(OUTPUT_DIR)/$(BINARY_NAME_CLI)$(EXT)

.PHONY: production/pack-all
production/pack-all: production/pack-webserver production/pack-cli ## Build and pack both binaries for local OS/ARCH

.PHONY: production/pack/matrix
.DELETE_ON_ERROR:
production/pack/matrix: $(MATRIX_BINARY_PATHS) ## Cross-compile and pack all OS/ARCH combinations

# Pattern rule for matrix builds
$(OUTPUT_DIR)/%:
	$(call build_and_pack,$@)

