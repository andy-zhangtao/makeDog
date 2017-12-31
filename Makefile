
.PHONY: build
name = makeDog

build:
	go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d-%H%M%S)" -o $(name)

run: build
	./$(name)

release: *.go *.md
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d)" -a -o $(name)
