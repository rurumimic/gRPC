# productinfo

- [ProductInfo Service and Client](https://github.com/grpc-up-and-running/samples/tree/master/ch02)

---

## compile proto

### Server

in `./productinfo`:

```bash
python -m grpc_tools.protoc \
       -Iproto \
       --python_out=python/server \
       --grpc_python_out=python/server \
       proto/product_info.proto
```

### Client

in `./productinfo`:

```bash
python -m grpc_tools.protoc \
       -Iproto \
       --python_out=python/client \
       --grpc_python_out=python/client \
       proto/product_info.proto
```

### output

- [server/product_info_pb2.py](server/product_info_pb2.py)
- [server/product_info_pb2_grpc.py](server/product_info_pb2_grpc.py)
- [client/product_info_pb2.py](client/product_info_pb2.py)
- [client/product_info_pb2_grpc.py](client/product_info_pb2_grpc.py)

```bash
productinfo
├── proto
│   └── product_info.proto
└── python
    ├── client
    │   ├── product_info_pb2.py
    │   └── product_info_pb2_grpc.py
    └── server
        ├── product_info_pb2.py
        └── product_info_pb2_grpc.py
```

---

## Stub

### Server

```bash
python/server
├── product_info_pb2.py
├── product_info_pb2_grpc.py
└── server.py
```

### Client

```bash
python/client
├── client.py
├── product_info_pb2.py
└── product_info_pb2_grpc.py
```

---

## Run

### Run Server

```bash
cd ./productinfo/python/server
python server.py
```

```bash
Starting server. Listening on port 50051.
```

### Run Client

```bash
cd ./productinfo/python/client
python client.py
```

### output

```bash
# bin/server

Starting server. Listening on port 50051.
addProduct:request id: "76891aee-fae0-11ec-834c-8c85907c063a"
name: "Apple iPhone 11"
description: "Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode."
price: 699.0

addProduct:response value: "76891aee-fae0-11ec-834c-8c85907c063a"

getProduct:request value: "76891aee-fae0-11ec-834c-8c85907c063a"

getProduct:response id: "76891aee-fae0-11ec-834c-8c85907c063a"
name: "Apple iPhone 11"
description: "Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode."
price: 699.0
```

```bash
# bin/client

add product: response value: "76891aee-fae0-11ec-834c-8c85907c063a"

get product: response id: "76891aee-fae0-11ec-834c-8c85907c063a"
name: "Apple iPhone 11"
description: "Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode."
price: 699.0
```
