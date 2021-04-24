# 1. Set up  the golang

#### Download go from official website

#### Set up the environment path

reminder: distinguish GOPATH  GOROOT GOBIN

GOPATH: go get : all dependencies will be in this directory

GOROOT: where you install the go

GOBIN: GOROOT / bin

:last_quarter_moon_with_face: one of the most import thing is do not forget to add BIN path to local path variable

# 2 Install the proto buffer

### step1. go to the github to get the protobuf

Link: https://github.com/google/protobuf/releases
version:3.15.8


download the resource then move the protoc.exe file to the GOBIN directory

### step2. set up the plugins

1. go get github.com/golang/protobuf
2. go get -u github.com/golang/protobuf/protoc-gen-go
3. go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

### step2. write the .proto

​	run the command

```bash
protoc greet/greetpb/greet.proto --go_out=. --go-grpc_out=.
```

## 3. Problem

1. can not find package

   solutionist clone to the corresponding directory
