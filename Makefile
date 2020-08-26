help:
	@echo "Commands:"
	@echo "make run - run app in docker"
	@echo "make tests-unit - run unit tests"
	@echo "make tests-e2e - run integration tests"
	@echo "make dev-run - run in dev mod"

run:
	docker-compose up -d --build

test-units:
	go test -race -count 100 ./... -v

test-e2e:
	go test -tags e2e ./... -v -count=1 -parallel=1


migrate-create:
	docker-compose run --no-deps migrate create -ext sql migration

migrate-up:
	docker-compose run --no-deps migrate up


lint:
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.30.0 golangci-lint --color always run -v

lint-fix:
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.30.0 golangci-lint --color alwaay run -v --fix ./..

ci-build:
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
