# productinfo client

## mod

```bash
go mod init productinfo/client
```

```bash
go: creating new go.mod: module productinfo/client
```

### output

#### go.mod

```go
module productinfo/client

go 1.18
```

---

## main.go

Write: [main.go](main.go)

---

## auto-import dependencies

```bash
go mod tidy
```

```bash
go: finding module for package google.golang.org/protobuf/runtime/protoimpl
go: finding module for package google.golang.org/protobuf/reflect/protoreflect
go: finding module for package google.golang.org/grpc/codes
go: finding module for package google.golang.org/grpc
go: finding module for package google.golang.org/grpc/status
go: found google.golang.org/grpc in google.golang.org/grpc v1.47.0
go: found google.golang.org/grpc/codes in google.golang.org/grpc v1.47.0
go: found google.golang.org/grpc/status in google.golang.org/grpc v1.47.0
go: found google.golang.org/protobuf/reflect/protoreflect in google.golang.org/protobuf v1.28.0
go: found google.golang.org/protobuf/runtime/protoimpl in google.golang.org/protobuf v1.28.0
```

### output

- go.mod
- go.sum

#### go.mod

```go
module productinfo/client

go 1.18

require (
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
```
