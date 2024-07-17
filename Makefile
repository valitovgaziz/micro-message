build:
	@go build -o bin/fs

run: build
	@./bin/fs
test:
	@go test

.DEFAULT_GOAL=run