# Helloworld

- [Quick start](https://grpc.io/docs/languages/go/quickstart/): Helloworld
- example code: [grpc/grpc-go/helloworld](https://github.com/grpc/grpc-go/tree/master/examples/helloworld)
  - pkg: [google.golang.org/grpc/examples/helloworld/helloworld](https://pkg.go.dev/google.golang.org/grpc/examples/helloworld/helloworld)
  - [helloworld/helloworld](https://github.com/grpc/grpc-go/tree/master/examples/helloworld/helloworld)
    - helloworld.pb.go
    - helloworld.proto
    - helloworld_grpc.pb.go

## Init

```bash
go mod init helloworld # create go.mod
go mod tidy # generate go.sum and update go.mod
```

## Run

Run server: termianl 1

```bash
go run greeter_server/main.go

2022/07/02 19:59:38 server listening at 127.0.0.1:50051
```

Run client: termianl 2

```bash
go run greeter_client/main.go

2022/07/02 19:59:50 Greeting: Hello world
```

```bash
go run greeter_client/main.go --name=Alice

2022/07/03 15:25:35 Greeting: Hello Alice
```

termianl 1:

```bash
go run greeter_server/main.go

2022/07/02 19:59:38 server listening at 127.0.0.1:50051
2022/07/02 19:59:50 Received: world
2022/07/03 15:25:35 Received: Alice
```
