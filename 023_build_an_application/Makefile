.DEFAULT_GOAL := help

# Change these variables as necessary
MAIN_PACKAGE_PATH := .
BINARY_NAME := http-server
OUTPUT_DIR := .

# Set the OS and ARCH options based for cross-compilation
OS_LIST := linux darwin windows
ARCH_LIST := amd64 arm64

MATRIX_BINARY_NAMES := $(foreach os,$(OS_LIST),$(foreach arch,$(ARCH_LIST),$(BINARY_NAME)_$(os)_$(arch)))
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
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
	
.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	@if [ -n "$(shell git status --porcelain)" ]; then \
		echo "Error: Working directory is not clean. Please commit or stash your changes."; \
		exit 1; \
	fi

.PHONY: coverage
coverage:
	mkdir -p $@

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
.PHONY: audit
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)" 
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs=true ./...

## test/cover: run all tests and display coverage
.PHONY: test/coverage
test/cover: coverage
	go test -v -race -buildvcs=true -coverprofile=${OUTPUT_DIR}/coverage/coverage.out ./...
	go tool cover -html=${OUTPUT_DIR}/coverage/coverage.out -o ${OUTPUT_DIR}/coverage/coverage.html

## upgradeable: list direct dependencies that have upgrades available
.PHONY: upgradeable
upgradeable:
	@go list -u -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -m all


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

## build: build the application
.PHONY: build
build:
	GOOS=${OS} GOARCH=${ARCH} go build -o=${OUTPUT_DIR}/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the  application
.PHONY: run
run: build
	${OUTPUT_DIR}/bin/${BINARY_NAME}

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@latest \
		--build.cmd "make build" --build.bin "${OUTPUT_DIR}/bin/${BINARY_NAME}" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"


# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

## push: push changes to the remote Git repository
.PHONY: push
push: confirm audit no-dirty
	git push

## production/deploy: deploy the application to production
.PHONY: production/deploy
production/deploy: audit no-dirty
	GOOS=${OS} GOARCH=${ARCH} go build -ldflags='-s' -o=${OUTPUT_DIR}/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

	@if [ "${OS}" = "darwin" ]; then \
		upx --best --lzma --force-macos ${OUTPUT_DIR}/bin/${BINARY_NAME} || (echo "> \033[31m\033[1mUPX failed for ${OS}/${ARCH}, skipping...\033[0m"; true); \
	else \
		upx --best --lzma ${OUTPUT_DIR}/bin/${BINARY_NAME} || (echo "> \033[31m\033[1mUPX failed for ${OS}/${ARCH}, skipping...\033[0m"; true); \
	fi
