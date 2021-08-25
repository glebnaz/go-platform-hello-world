.PHONY: generate
generate: generate-proto-go

.PHONY: generate-proto-go
generate-proto-go:
	buf mod update
	buf build
	buf generate -v --path=api/v1
	# add swagger.json to bindata
	go-bindata -nometadata --nocompress -pkg bindata -o internal/pkg/bindata/swagger-json.go api/api.swagger.json
