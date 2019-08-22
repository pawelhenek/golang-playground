Golang playground
===

![Circle CI build](https://circleci.com/gh/pawelhenek/golang-playground/tree/master.svg?style=svg)

#### Prerequisites:
- Installed Go 11.1 (https://golang.org/doc/install#install)

#### Quick start

`git clone https://github.com/pawelhenek/golang-playground.git`

`cd golang-playground`

Running snippets from https://gobyexample.com

`go run src/github.com/pawelhenek/go-by-example/for.go`

Running examples from https://gowebexamples.com

`go run src/github.com/pawelhenek/go-web-examples/hello-server.go`

To run examples from https://quii.gitbook.io (when `$GOPATH` is set to e.g. `/home/pawelhenek/Pulpit/golang-playground/`) 

`src/github.com/pawelhenek/golang-playground/learn-go-with-tests/a-working-application` 

`go test -v`  

#### Locally verified with

`go version`

`go version go1.11.4 linux/amd64`

#### Try out 
- configure tests execution in pipeline
- more go syntax & go modules examples
- go build & dependencies tools
- testing libs
- web/http/api frameworks
- dbs, msging, protocols libs