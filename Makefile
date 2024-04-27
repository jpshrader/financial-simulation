run:
	go run main.go

test:
	go test ./...

build:
	go build -v ./...

update:
	go get -u
	go mod tidy