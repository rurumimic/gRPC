# gRPC

- [gRPC](https://grpc.io/)
  - [documentation](https://grpc.io/docs/)
  - [introduction](https://grpc.io/docs/what-is-grpc/introduction/)
  - quickstart: [golang](https://grpc.io/docs/languages/go/quickstart/)
- github: [grpc/grpc-go](https://github.com/grpc/grpc-go)
- protocol buffers: [google doc](https://developers.google.com/protocol-buffers)

## Ref

- book: [gRPC - Up and Running](https://grpc-up-and-running.github.io/)
  - [github](https://github.com/grpc-up-and-running)
    - [content](https://github.com/grpc-up-and-running/grpc-up-and-running.github.io)
    - [sampels](https://github.com/grpc-up-and-running/samples)

---

Read [Documentation](docs/README.md)

---

## Install

### Go

- Go
  - [gvm](https://github.com/moovweb/gvm)
- Protocol buffer 3+
  - [install](https://grpc.io/docs/protoc-installation/)

```bash
go version # go version go1.18 darwin/amd64
protoc --version # libprotoc 3.19.4
```

### Go plugins for the protocol compiler

- [protocolbuffers/protobuf-go/cmd/protoc-gen-go](https://github.com/protocolbuffers/protobuf-go/tree/master/cmd/protoc-gen-go)
- [grpc/grpc-go/cmd/protoc-gen-go-grpc](https://github.com/grpc/grpc-go/tree/master/cmd/protoc-gen-go-grpc)

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```bash
export PATH="$PATH:$(go env GOPATH)/bin"

echo $PATH | grep -E "$(go env GOPATH)/bin"
```
