LDFLAGS:=-X 'gitlab.ae-rus.net/platform-go/version.serviceName=${PROJECT_NAME}'\
		 -X 'gitlab.ae-rus.net/platform-go/version.branchName=$(GIT_BRANCH)'\
		 -X 'gitlab.ae-rus.net/platform-go/version.commitSHA=$(GIT_HASH)'\
		 -X 'gitlab.ae-rus.net/platform-go/version.commitMessage=$(COMMIT_MESSAGE)'\
		 -X 'gitlab.ae-rus.net/platform-go/version.buildTime=$(BUILD_TS)'

.PHONY: generate
generate: generate-proto-go

.PHONY: generate-proto-go
generate-proto-go:
	# update buf mod
	buf mod update
	# buf build
	buf build
	# generate proto
	buf generate -v --path=api/v1
	# swagger
	go-bindata -nometadata --nocompress -pkg bindata -o internal/pkg/bindata/swagger-json.go api/api.swagger.json

.PHONY: build
build:
	CGO_ENABLED=0 go build -v -ldflags "$(LDFLAGS)" -o bin/local ./cmd

.PHONY: lint
lint:
	golangci-lint run --fix ./...
