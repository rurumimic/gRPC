# Echo Server

## Build

```bash
uv sync
source .venv/bin/activate
```

## Run

### gRPC Server

```bash
(.venv) python -m echo
```

Run: [echo_client](../echo_client/README.md)

```bash
Starting server. Listening on port 50051.
sendMessage:request title: "Hello?"
content: "Hello, Server!"

sendMessage:response title: "Hello?"
```

### REST API

```bash
(.venv) fastapi run
```

```bash
curl localhost:8000

{"Hello":"World"}
```

```bash
curl 'localhost:8000/items/1?q=foo'

{"item_id":1,"q":"foo"}
```

---

## Compile Protobuf

### Server

in `./proxy`

```bash
python -m grpc_tools.protoc \
       -Iproto \
       --python_out=python/echo/app/echo/rpc/message \
       --grpc_python_out=python/echo/app/echo/rpc/message \
       --pyi_out=python/echo/app/echo/rpc/message \
       proto/message.proto
```

edit:

```py
- import message_pb2 as message__pb2
+ from echo.rpc.message import message_pb2 as message__pb2
```

