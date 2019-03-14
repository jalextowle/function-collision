clean: 
	rm checker
build:
	go build -o checker
test:
	go test ./...
