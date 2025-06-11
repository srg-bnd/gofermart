TEST?=$$(go list ./... | grep -v 'vendor' | grep -v 'env' | grep -v 'utils')
NAME=ya41-56
VERSION=0.0.1
BINARY_SERVER=server
BINARY_ACCRUAL=accrual
OS_ARCH=darwin_amd64

default: build

build: build_server build

combine:
	(find cmd/gophermart/main.go; find internal/gophermart -type f -name "*.go") \
		| while read file; do echo "// ===== $file ====="; cat "$file"; echo ""; done > combined.go

build_accrual:
	go build -o ${BINARY_ACCRUAL} cmd/accrual/main.go

build_server:
	go build -o ${BINARY_SERVER} cmd/gophermart/main.go

test_server:
	go test ./... -timeout 60m --tags=server -v

test_accrual:
	go test ./... -timeout 60m --tags=accrual -v

run_test_mart:
	./gophermarttest \
		-test.v -test.run=^TestGophermart$ \
		-gophermart-binary-path=./server \
		-gophermart-host=localhost \
		-gophermart-port=8080 \
		-gophermart-database-uri="postgres://myuser:mypassword@localhost:5432/mydb" \
		-accrual-binary-path=./accrual \
		-accrual-host=localhost \
		-accrual-port=41222 \
		-accrual-database-uri="postgres://myuser:mypassword@localhost:5432/mydb"

lint:
	golangci-lint run ./...