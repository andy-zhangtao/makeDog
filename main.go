package main

import (
	"flag"
	"fmt"
	"text/template"
	"os"
)

const makefile = `
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

var _VERSION_ = "unknown"

var name = flag.String("name", "", "The binary name")

func main() {
	fmt.Println(getVersion())

	flag.Parse()

	if *name == "" {
		fmt.Println("Name can't be empty! See MakeDog Usage below\n")
		flag.Usage()
		os.Exit(-1)
	}

	type MakeFile struct {
		Name string
	}

	var mf = MakeFile{Name: *name}

	t := template.Must(template.New("makefile").Parse(makefile))

	file, err := os.Create("Makefile")
	if err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}

	err = t.Execute(file, mf)
	if err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println("MakeDog Finish! Wang~ Wang~")
}

func getVersion() string {
	return fmt.Sprintf("== MakeDog [%s] == \n", _VERSION_)
}
