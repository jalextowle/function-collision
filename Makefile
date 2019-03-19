clean: 
	rm checker
build:
	go build -o checker
run:
	./checker ./test_data/AppProxyUpgradeable.json ./test_data/ACL.json
test:
	go test ./...
