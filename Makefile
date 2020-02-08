BINARY_NAME=server

configure:
	go mod tidy
test:
	go test -race github.com/hobord/golang-poc-rest/...
run:
	go run main.go
clean:
	go clean
	rm -Rf build
build: configure
	go build -o build/$(BINARY_NAME) -v