# ================================================================= #
# HELPER
# ================================================================= #

## help: print thid help message
.PHONY: help
help:
	@echo 'Usage'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ================================================================= #
# DEVELOPMENT
# ================================================================= #

## run/app: run the 'cmd/app' application
.PHONY: run/app
run/app:
	go run ./cmd/app $(filter-out $@,$(MAKECMDGOALS))

# ================================================================= #
# QUALITY CONTROL
# ================================================================= #

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	# staticcheck ./....
	@echo 'Running tests...'
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ================================================================= #
# BUILD
# ================================================================= #

current_time = $(shell date --iso-8601=seconds)
git_description = $(shell git describe --always --tags)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

## build/app: build the 'cmd/app' application
.PHONY: build/app
build/app:
	@echo 'Building cmd/app...'
	go build -ldflags=${linker_flags} -o=./bin/xecute ./cmd/app
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/xecute_linux_amd64 ./cmd/app

# ================================================================= #
# TAGGING & RELEASE
# ================================================================= #

# Tag name can be passed like: make tag version=v1.0.0
version ?= $(shell git describe --tags --abbrev=0)

## tag: create a git tag
.PHONY: tag
tag:
	@if git rev-parse $(version) > /dev/null 2>&1; then \
		echo "Tag $(version) already exists! Aborting." \
		exit 1; \
	fi
	@git tag -a $(version) -m "Relase $(version)"
	@git push origin $(version)

## release: creates tag and builds 'cmd/app'
.PHONY: release
release: tag build/app

