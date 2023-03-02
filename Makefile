.PHONY: all
all: build
FORCE: ;

.PHONY: build

generate:
	weaver generate ./...

build: generate
	go build -o bin/microservices

run-single: build
	go run .

run-multi: build 
	weaver multi deploy weaver.toml

status:
	weaver multi status

dashboard:
	weaver multi dashboard