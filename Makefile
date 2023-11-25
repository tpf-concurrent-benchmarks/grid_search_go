
test_interval_available:
	cd ./src/manager/src/interval && go test -run TestAvailable

test_interval:
	cd ./src/manager/src/interval && go test

format:
	go fmt ./...
