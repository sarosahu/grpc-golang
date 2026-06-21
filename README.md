# grpc-golang

grpc in Go

https://grpc.io/docs/languages/go/quickstart/

Prerequisites

> > > > > > > > > > > > > Go, any one of the two latest major releases of Go.

For installation instructions, see Go’s Getting Started guide.

Protocol buffer compiler, protoc, version 3.

For installation instructions, see Protocol Buffer Compiler Installation.

Go plugins for the protocol compiler:

Install the protocol compiler plugins for Go using the following commands:

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
Update your PATH so that the protoc compiler can find the plugins:

export PATH="$PATH:$(go env GOPATH)/bin"

<<<<<<<<<<<<<<

Write proto file and generate code

> > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

1. Create a folder called greet and add 3 sub folders client, server, proto under greet/
2. Under greet/proto, add a proto file dummy.proto
3. Then run the following command:
   protoc -Igreet/proto --go_out=. --go_opt=module=github.com/sarosahu/grpc-golang --go-grpc_out=. --go-grpc_opt=module=github.com/sarosahu/grpc-golang greet/proto/dummy.proto

   Note: This will generate the source code files under proto. The generated files will be named as dummy_grpc.pb.go and dummy.pb.go

4. Now instead of executing this long command (step 3), we can have a Makefile and run it as follows:
   make greet

   Note: This will generate the same \*.go files under proto folder.
   Here is the snapshot of the output:

   > > > make greet
   > > > protoc -Igreet/proto --go_opt=module=github.com/sarosahu/grpc-golang --go_out=. --go-grpc_opt=module=github.com/sarosahu/grpc-golang --go-grpc_out=. greet/proto/\*.proto
   > > > go build -o bin/greet/server ./greet/server
   > > > no Go files in /Users/sarojkumarsahu/ml/mygithub/grpc/grpc-golang/greet/server
   > > > make: \*\*\* [greet] Error 1

   Note: The error is expected.

5. We can run 'make help' to list all the make commands for this project, here is the output of 'make help':
   make help
   all Generate Pbs and build
   greet Generate Pbs and build for greet
   calculator Generate Pbs and build for calculator
   blog Generate Pbs and build for blog
   test Launch tests
   clean Clean generated files
   clean_greet Clean generated files for greet
   clean_calculator Clean generated files for calculator
   clean_blog Clean generated files for blog
   rebuild Rebuild the whole project
   bump Update packages version
   about Display info related to the build
   help Show this help

   make clean_greet
   rm -f greet/proto/\*.pb.go

   make about
   OS: macos 25.4.0 arm64
   Shell: bash 3.2.57(1)-release
   Protoc version: libprotoc 35.0
   Go version: go version go1.26.0 darwin/arm64
   Go package: github.com/sarosahu/grpc-golang
   Openssl version: OpenSSL 3.5.1 1 Jul 2025 (Library: OpenSSL 3.5.1 1 Jul 2025)

6.
