# Continuous Learning: Golang

Learning Golang for Ops and Site Reliability Engineering. 

## Why Go for Ops and SRE? 

* Open Source
* Simple, reliable and efficient
* Testing, benchmarking and profiling built-in
* Good standard libraries
* Concurrency via goroutines and channels
* Google's definition SRE is a software engineer running operations
* SREs normally write tools and projects which are more complex than 'scripts' - let's eliminate real toils! :P
* Go in production at Google and outside!

### Go Modules

* Aim to solve problems related to dependency management, version selection and reproducible builds
* Enable users to run Go code outside of GOPATH

### Go Doc
* To install run: $ go get golang.org/x/tools/cmd/godoc
* Launch the docs locally by running: $ godoc -http :8000
* http://localhost:8000/pkg: see all the packages installed

### Go Linting: golangci-lint
```bash
# Ubuntu 
sudo snap install golangci-lint --edge
golangci-lint run
```

## Learning Resources and Credits

* https://github.com/quii/learn-go-with-tests
* https://awesome-go.com/
* https://github.com/golang/vscode-go/blob/master/docs/tools.md
* https://github.com/go-modules-by-example/index/blob/master/009_submodules/README.md
