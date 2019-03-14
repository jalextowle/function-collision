TEST1_SELECTOR="0x00000000"
TEST1_ABI_PATH="./test/abi.txt"

clean: 
	rm checker
build:
	go build -o checker
test:
	make build
	./checker ${TEST1_SELECTOR} ${TEST1_ABI_PATH}
#run:
#	./checker 
