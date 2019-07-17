
.PHONY: build
name = makeDog
version = $(shell date +%Y%m%d-%H%M%S)

build:
	go build -ldflags "-X main._VERSION_=$(version)" -o $(name)

run: build
	./$(name)

release: *.go *.md
	git rev-parse --short HEAD~2|xargs git rev-list --format=%B --max-count=1|xargs echo `date`  > build.info
	docker run -it --rm --name golang -e GOOS=linux -e GOARCH=amd64 -v $(PWD):/go/src/dest -w /go/src/dest golang:1.12-alpine go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d)" -a -o bin/$(name)
	docker build -t vikings/$(name) .
	docker push vikings/$(name)
