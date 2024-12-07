# Echo Client

## Build

```bash
uv sync
source .venv/bin/activate
```

## Run

```bash
(.venv) python -m echo_client

server title: "Hello?"
```

---

## Compile Protobuf

### Client

in `./proxy`

```bash
python -m grpc_tools.protoc \
       -Iproto \
       --python_out=python/echo_client/rpc/message \
       --grpc_python_out=python/echo_client/rpc/message \
       --pyi_out=python/echo_client/rpc/message \
       proto/message.proto
```

edit:

```py
- import message_pb2 as message__pb2
+ from echo_client.rpc.message import message_pb2 as message__pb2
```

