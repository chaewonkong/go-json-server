# build
build:
	@mkdir -p bin
	@go mod download && go mod verify
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/
.PHONY: build

# install dependencies
ref:
	@go mod tidy -v
	@go mod vendor -v
.PHONY: ref