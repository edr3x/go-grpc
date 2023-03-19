# GRPC with go

## Prerequisites

- [GO](https://go.dev/doc/install)
- [Protocol buffer compiler](https://grpc.io/docs/protoc-installation/)
- Go plugins for the protocol compiler

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

- Update path
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Get necessary packages for project

```bash
go mod tidy
```

## Create `*.pb.go` file

```bash
protoc --go_out=. --go-grpc_out=. <relative-path-to-proto-file> # `proto/greet.progo` for this project
```

## Open two terminals, one on `server/` and another on `client/` directory and run

```bash
go run *.go # runs all go files on current directory
```
