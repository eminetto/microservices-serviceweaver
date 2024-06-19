.PHONY: all
all: build
FORCE: ;

.PHONY: build

generate:
	weaver generate ./...

build: generate
	go build -o bin/microservices

run-single: build
	SERVICEWEAVER_CONFIG=weaver.toml go run .

run-multi: build 
	weaver multi deploy weaver.toml

status:
	weaver multi status

dashboard:
	weaver multi dashboard

gke-run-multi: build
	weaver gke deploy weaver.toml

gke-status:
	weaver gke status

gke-dashboard:
	weaver gke dashboard

gke-local-run-multi: build
	weaver gke-local deploy weaver.toml

gke-local-status:
	weaver gke-local status	