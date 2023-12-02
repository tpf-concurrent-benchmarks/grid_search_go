N_WORKERS=2

init:
	docker swarm init

build:
	docker rmi grid_search_go_worker -f
	docker rmi grid_search_go_manager -f
	docker build -t grid_search_go_worker ./src/ -f ./src/worker/Dockerfile
	docker build -t grid_search_go_manager ./src/ -f ./src/manager/Dockerfile

setup: init build

remove:
	docker stack rm gs_go

deploy:
	mkdir -p graphite
	N_WORKERS=${N_WORKERS} docker compose -f=docker-compose-deploy-local.yml up --build

deploy_remote:
	mkdir -p graphite
	N_WORKERS=${N_WORKERS} docker stack deploy -c docker-compose-deploy.yml gs_go

down_graphite:
	if docker stack ls | grep -q graphite; then \
		docker stack rm graphite; \
		docker stack rm grafana; \
		docker stack rm cadvisor; \
	fi
.PHONY: down_graphite

test_interval_available:
	cd ./src/manager/src/interval && go test -run TestAvailable

test_interval:
	cd ./src/manager/src/interval && go test

run_worker_local:
	cd ./src/worker && LOCAL=local go run ./src

run_manager_local:
	cd ./src/manager && LOCAL=local go run ./src

format:
	cd ./src/common && go fmt ./...
	cd ./src/manager && go fmt ./...
	cd ./src/worker && go fmt ./...
