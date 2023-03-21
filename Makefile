.PHONY: generate build 

all: clean generate build

clean:
	rm -rf build/server
	go clean -i .

validate:
	swagger validate ./api/swagger.yml

doc: validate
	swagger serve api/swagger.yml --no-open --host=0.0.0.0 --port=8080 --base-path=/

generate: validate
	go generate ./...

run:
	./build/main start

run-local:
	go run cmd/server/main.go start

build:
	go build -ldflags="-w -s" -o build/main ./cmd/server 

migrate-up:
	go run cmd/server/main.go migrations

migrate-down:
	go run cmd/server/main.go rollback

migrate-drop:
	go run cmd/server/main.go drop

help:
	go run cmd/server/main.go -h