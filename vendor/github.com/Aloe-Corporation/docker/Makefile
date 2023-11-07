# env
project_name:=$(shell grep "module " <go.mod | sed 's;.*/;;g')

format:
	go fmt ./...

test:
	go clean -testcache
	go test -cover ./...

upd-vendor:
	go mod tidy
	go mod vendor

lint: 
	golangci-lint run --allow-parallel-runners -c ~/.golangci-lint.yml --fix ./...

.PHONY: format upd-vendor test lint