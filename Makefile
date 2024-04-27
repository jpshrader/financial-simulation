run:
	go run .

test:
	go test -cover ./...

update:
	go get -u
	go mod tidy