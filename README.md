# makeDog
A golang tool can auto create Makefile

## Usage

```$golang
Usage of makeDog:
  -name string
    	The binary name
```

## Makefile
The default makefile content is :
```$makefile
.PHONY: build
name = {{Name}}

build:
	go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d-%H%M%S)" -o $(name)

run: build
	./$(name)

release: *.go *.md
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d)" -a -o $(name)
	docker build -t vikings/$(name) .
	docker push vikings/$(name)

```
