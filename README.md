Golang playground
===

#### Prerequisites:
- Installed Go with modules support 1.19 (https://golang.org/doc/install#install)
- Installed [just](https://github.com/casey/just)

#### Quick start

`git clone https://github.com/pawelhenek/golang-playground.git &&  cd golang-playground`

##### With just

```bash
just -l
Available recipes:
    hello
    run-go-by-example
    run-learn-go-with-tests
```

Run all script from https://gobyexample.com

`just run-go-by-example`

Run all examples from https://gowebexamples.com

`go run src/github.com/pawelhenek/go-web-examples/hello-server.go`

To run examples from https://quii.gitbook.io (when `$GOPATH` is set to e.g. `/home/pawelhenek/Pulpit/golang-playground/`) 

`just run-go-by-example`

#### Locally verified with

`go version`
`go version go1.19 linux/amd64`

#### Try out 
- configure tests execution in pipeline
- more go syntax & go modules examples
- go build & dependencies tools
- testing libs
- web/http/api frameworks
- grpc
- samples with test containers
- k8s api
- terraform provider
- kafka
- postgresql
- cli app
- dagger pipelines
- docker, earthly build files