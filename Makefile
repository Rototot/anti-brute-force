help:
	@echo "Commands:"
	@echo "make init-dev - init dev env"
	@echo "make init-prod - init prod env"
	@echo "make run - run app"
	@echo "make install - install deps"
	@echo "make build - compile apps"
	@echo "make tests-unit - run unit tests"
	@echo "make tests-e2e - run e2e tests"

run: build
	./build/web

build: install
	rm -r -f bin/*
	go build -i -o build/web ./cmd/web

install:
	go mod download

tests-unit: install
	go test -race -count 100 ./... -v

tests-e2e: install
	go test -tags e2e ./... -v -count=1 -parallel=1


init-dev: clear-init
	cp .env.dist .env
	cp docker-compose.override.yml.dist docker-compose.override.yml

init-prod: clear-init
	cp .env.dist .env

clear-init:
	rm .env docker-compose.override.yml