help:
	@echo "Commands:"
	@echo "make run - run app in docker"
	@echo "make tests-unit - run unit tests"
	@echo "make tests-e2e - run integration tests"
	@echo "make dev-run - run in dev mod"

run:
	docker-compose up -d --build

tests-unit:
	go test -race -count 100 ./... -v

tests-e2e:
	go test -tags e2e ./... -v -count=1 -parallel=1

dev-run: dev-build
	./build/app

dev-build: init-dev
	rm -r -f build/*
	go build -i -o build/app .

#dev-install:
	#go mod download

init-dev: clear-init
	cp .env.dist .env
	cp docker-compose.override.yml.dist docker-compose.override.yml

init-prod: clear-init
	cp .env.dist .env

clear-init:
	rm .env docker-compose.override.yml
