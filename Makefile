BINARY_NAME=api-template-go

build:
	go build -o ${BINARY_NAME}.exe main.go

run:
	./${BINARY_NAME}.exe

build_and_run: build run

clean:
	go clean

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet