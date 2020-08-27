help:
	@echo "Commands:"
	@echo "make build - build app in docker"
	@echo "make run - run app"
	@echo "make test - run all tests"


build:
	docker-compose build

run: build migrate-up
	docker-compose up -d

test: migrate-up-test test-units test-e2e
	docker-compose run -f docker-compose.test.yaml --env-file .env.test --entrypoin "" app make test-units
	docker-compose run -f docker-compose.test.yaml --env-file .env.test --entrypoin "" app make test-e2e

test-units:
	go test -race -count 100 ./... -v

test-e2e:
	go test -tags e2e ./... -v -count=1 -parallel=1


migrate-create:
	docker-compose run --no-deps migrate create -ext sql migration

migrate-up:
	docker-compose run migrate up

migrate-up-test:
	docker-compose -f docker-compose.test.yaml  --env-file .env.test run --no-deps migrate up

lint:
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.30.0 golangci-lint --color always run -v

lint-fix:
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.30.0 golangci-lint --color alwaay run -v --fix ./..

ci-build:
	go generate
	rm -r -f bin/*
	go build -i -o bin/app .

dev-run: dev-build
	./bin/app

dev-build: init-dev
	rm -r -f bin/*
	go build -i -o bin/app .

#dev-install:
	#go mod download

init-dev: clear-init
	cp .env.dist .env
	cp docker-compose.override.yml.dist docker-compose.override.yml

init-prod: clear-init
	cp .env.dist .env

clear-init:
	rm .env docker-compose.override.yml
