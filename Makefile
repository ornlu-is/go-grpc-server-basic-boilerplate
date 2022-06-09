SHELL=/bin/bash
TEST_COVERAGE_THRESHOLD=63.8

.SILENT: proto test lint docs build run

local-test: proto test lint docs

proto:
	protoc api.proto --go_out=. \
    				 --go_opt=paths=source_relative \
    				 --go-grpc_out=. \
    				 --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
    				 --proto_path=. 

test:
	cd server; go test -coverprofile=../coverage.out 
	coverage=$$(go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+') ;\
	rm coverage.out ;\
	if [ $$(bc <<< "$$coverage < $(TEST_COVERAGE_THRESHOLD)") -eq 1 ]; then \
		echo "Low coverage: $$coverage < $(TEST_COVERAGE_THRESHOLD)" ;\
		exit 1 ;\
	fi ;\

lint:
	golangci-lint run --config ./.golangci.yaml

docs:
	protoc --doc_out=. --doc_opt=template.tmpl,docs.md api.proto

build:
	cd server; go build -o server; mv server ../bin

run: 
	./bin/server