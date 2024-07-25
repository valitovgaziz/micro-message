build:
	@go build -o bin/fs

run: build
	@./bin/fs
test:
	@go test

dbuild:
	@docker compose build

dup: dbuild
	@docker compose up

.DEFAULT_GOAL=dup