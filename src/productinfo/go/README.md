# productinfo

- [ProductInfo Service and Client](https://github.com/grpc-up-and-running/samples/tree/master/ch02)

---

## compile proto

### Server

in `./productinfo`:

```bash
protoc -I proto \
       proto/product_info.proto \
       --go_out=go/server \
       --go-grpc_out=go/server
```

OR

in `./productinfo/go/server`:

```bash
protoc -I ../../proto \
       ../../proto/product_info.proto \
       --go_out=. \
       --go-grpc_out=.
```

### Client

in `./productinfo`:

```bash
protoc -I proto \
       proto/product_info.proto \
       --go_out=go/client \
       --go-grpc_out=go/client
```

### output

- [server/ecommerce/product_info.pb.go](server/ecommerce/product_info.pb.go)
- [server/ecommerce/product_info_grpc.pb.go](server/ecommerce/product_info_grpc.pb.go)
- [client/ecommerce/product_info.pb.go](client/ecommerce/product_info.pb.go)
- [client/ecommerce/product_info_grpc.pb.go](client/ecommerce/product_info_grpc.pb.go)

```bash
productinfo
├── go
│   ├── client
│   │   └── ecommerce
│   │       ├── product_info.pb.go
│   │       └── product_info_grpc.pb.go
│   └── server
│       └── ecommerce
│           ├── product_info.pb.go
│           └── product_info_grpc.pb.go
└── proto
    └── product_info.proto
```

---

## Stub

### Server

- [server/README](server/README.md)

```bash
productinfo/go/server
├── README.md
├── ecommerce # stub files
│   ├── product_info.pb.go
│   └── product_info_grpc.pb.go
├── go.mod
├── go.sum
└── main.go
```

### Client

- [client/README](client/README.md)

```bash
productinfo/go/client
├── README.md
├── ecommerce # stub files
│   ├── product_info.pb.go
│   └── product_info_grpc.pb.go
├── go.mod
├── go.sum
└── main.go
```

---

## Build

### Build Server

```bash
cd ./productinfo/go/server
go build -v -o bin/server
```

#### output

```bash
productinfo/go/server
└── bin
   └── server
```

### Build Client

```bash
cd ./productinfo/go/client
go build -v -o bin/client
```

```bash
productinfo/go/client
└── bin
    └── client
```

---

## Run

### Run Server

```bash
cd ./productinfo/go/server
bin/server
```

### Run Client

```bash
cd ./productinfo/go/client
bin/client
```

### output

```bash
# bin/server

2022/07/03 17:46:23 Product c15194bc-66b4-4d01-9355-61ebca00695d : Apple iPhone 11 - Added.
2022/07/03 17:46:23 Product c15194bc-66b4-4d01-9355-61ebca00695d : Apple iPhone 11 - Retrieved.
```

```bash
# bin/client

2022/07/03 17:46:23 Product ID: c15194bc-66b4-4d01-9355-61ebca00695d added successfully
2022/07/03 17:46:23 Product: id:"c15194bc-66b4-4d01-9355-61ebca00695d" name:"Apple iPhone 11" description:"Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode." price:699
```
