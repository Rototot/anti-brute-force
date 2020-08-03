help:
	@echo "Commands:"
	@echo "make install - install deps"
	@echo "make build - compile apps"
	@echo "make run - run app"
	@echo "make tests-unit - run unit tests"
	@echo "make tests-e2e - run e2e tests"

run: build
	./build/web

build: install
	go build ./cmd bi
	rm -r -f bin/*
	go build -i -o build/web ./cmd/web

install:
	go mod download

tests-unit:
	go test -race -count 100 ./... -v

tests-e2e:
	go test -tags e2e ./... -v -count=1 -parallel=1
