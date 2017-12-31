package main

import (
	"flag"
	"fmt"
	"text/template"
	"os"
)

const makefileWithDocker = `
.PHONY: build
name = {{.Name}}

build:
	go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d-%H%M%S)" -o $(name)

run: build
	./$(name)

release: *.go *.md
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d)" -a -o $(name)
	docker build -t vikings/$(name) .
	docker push vikings/$(name)
`
const makefileWithBinary = `
.PHONY: build
name = {{.Name}}

build:
	go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d-%H%M%S)" -o $(name)

run: build
	./$(name)

release: *.go *.md
	GOOS={{.Release}} GOARCH=amd64 go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d)" -a -o $(name)
`
const (
	DOCKER = "docker"
	BINARY = "binary"
	LINUX  = "linux"
	MACOS  = "macos"
	WINDOW = "windows"
)

var _VERSION_ = "unknown"

var name = flag.String("name", "", "The binary name")
var kind = flag.String("kind", DOCKER, "Release A Docker Image")
var releaseOS = flag.String("release_os", LINUX, "The Release OS Name. linux/darwin/windows")
var version = flag.Bool("version", false, "Output MakeDog Version")

func main() {

	flag.Parse()

	if *version {
		fmt.Println(getVersion())
		os.Exit(0)
	}
	
	if *name == "" {
		fmt.Println("Name can't be empty! See MakeDog Usage below\n")
		flag.Usage()
		os.Exit(-1)
	}

	var t *template.Template
	switch *kind {
	case DOCKER:
		t = template.Must(template.New("makefile").Parse(makefileWithDocker))
	case BINARY:
		t = template.Must(template.New("makefile").Parse(makefileWithBinary))
	}

	type MakeFile struct {
		Name    string
		Release string
	}

	var mf = MakeFile{
		Name:    *name,
		Release: *releaseOS,
	}

	file, err := os.Create("Makefile")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	err = t.Execute(file, mf)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println("MakeDog Finish! Wang~ Wang~")
}

func getVersion() string {
	return fmt.Sprintf("== MakeDog [%s] == \n", _VERSION_)
}
