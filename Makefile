
test_interval_available:
	cd ./src/manager/src/interval && go test -run TestAvailable

test_interval:
	cd ./src/manager/src/interval && go test

run_worker_local:
	cd ./src/worker && ENV=local go run ./src

run_manager_local:
	cd ./src/manager && ENV=local go run ./src

format:
	go fmt ./...
