TEST?=$$(go list ./... | grep -v 'vendor' | grep -v 'env' | grep -v 'utils')
NAME=ya41-56
VERSION=0.0.1
BINARY_SERVER=server
BINARY_ACCRUAL=accrual
OS_ARCH=darwin_amd64
COMBINED_FILE := combined.go

default: build

build: build_server build

combine:
	@echo "Combining Go files into $(COMBINED_FILE)..."
	@rm -f $(COMBINED_FILE)
	@{ \
		echo cmd/flags.go; \
		find cmd/accrual -type f -name "*.go"; \
		find internal/accrual -type f -name "*.go"; \
		find internal/shared -type f -name "*.go"; \
	} | while read -r file; do \
		echo "// ===== $$file =====" >> $(COMBINED_FILE); \
		cat "$$file" >> $(COMBINED_FILE); \
		echo "" >> $(COMBINED_FILE); \
	done
	@echo "Done: $(COMBINED_FILE)"

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