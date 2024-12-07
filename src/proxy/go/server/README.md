# Server

---
## Compile Protobuf

### Server

in `./proxy`

```bash
protoc -I proto \
       --go_out=go/server/rpc/message \
       --go_opt=paths=source_relative \
       --go-grpc_out=go/server/rpc/message \
       --go-grpc_opt=paths=source_relative \
       proto/message.proto
```

