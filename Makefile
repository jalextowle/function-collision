clean: 
	rm checker
build:
	go build -o checker
run:
	./checker "0x00000000" ./abi_data/ABI.json
test:
	go test ./...
