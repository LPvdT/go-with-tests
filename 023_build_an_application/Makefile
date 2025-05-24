.DEFAULT_GOAL := help

# Change these variables as necessary.
main_package_path := .
binary_name := http-server
output_dir := .

# Set the OS and architecture
OS := linux
ARCH := amd64

# Change the binary name depending on the OS
ifeq "$(OS)" "windows"
	binary_name := ${binary_name}_${OS}_${ARCH}.exe
endif
ifeq "$(OS)" "darwin"
	binary_name := ${binary_name}_${OS}_${ARCH}
endif
ifeq "$(OS)" "linux"
	binary_name := ${binary_name}_${OS}_${ARCH}
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
	@test -z "$(shell git status --porcelain)"

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
	go test -v -race -buildvcs=true -coverprofile=${output_dir}/coverage/coverage.out ./...
	go tool cover -html=${output_dir}/coverage/coverage.out -o ${output_dir}/coverage/coverage.html

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
	GOOS=${OS} GOARCH=${ARCH} go build -o=${output_dir}/bin/${binary_name} ${main_package_path}

## run: run the  application
.PHONY: run
run: build
	${output_dir}/bin/${binary_name}

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@latest \
		--build.cmd "make build" --build.bin "${output_dir}/bin/${binary_name}" --build.delay "100" \
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
production/deploy: confirm audit no-dirty
	GOOS=${OS} GOARCH=${ARCH} go build -ldflags='-s' -o=${output_dir}/bin/linux_amd64/${binary_name} ${main_package_path}
	upx -5 ${output_dir}/bin/linux_amd64/${binary_name}
