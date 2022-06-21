BINARY_NAME=api-template-go
EXEC_DIRECTORY=cmd/api-template-go

build:
	go build -o ${EXEC_DIRECTORY}/${BINARY_NAME}.exe ${EXEC_DIRECTORY}/main.go

run:
	./${EXEC_DIRECTORY}/${BINARY_NAME}.exe

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
	go vet ./...