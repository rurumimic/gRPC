# productinfo

- [ProductInfo Service and Client](https://github.com/grpc-up-and-running/samples/tree/master/ch02)

## protobuf

- [proto/product_info.proto](proto/product_info.proto)

### syntax version

```proto
syntax = "proto3";
```

### package name

- unique names based on the project name
- prevent name clashes between protocol message types

```proto
package ecommerce;
```

### custom message

```proto
message ProductId {
  string value = 1;
}
```

```proto
message Product {
  string id = 1;
  string name = 2;
  string description = 3;
  float price = 4;
}
```

### service interface

= remote methods

```proto
service ProductInfo {
    rpc addProduct(Product) returns (ProductID);
    rpc getProduct(ProductID) returns (Product);
}
```

### option go_package

```proto
option go_package = "./ecommerce";
```


### (example) import proto

```proto
syntax = "proto3";

import "google/protobuf/wrappers.proto";

package ecommerce;
```

### compile proto

#### Server

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

#### Client

in `./productinfo`:

```bash
protoc -I proto \
       proto/product_info.proto \
       --go_out=go/client \
       --go-grpc_out=go/client
```

#### output

- [go/server/ecommerce/product_info.pb.go](go/server/ecommerce/product_info.pb.go)
- [go/server/ecommerce/product_info_grpc.pb.go](go/server/ecommerce/product_info_grpc.pb.go)

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

- [go/server/README](go/server/README.md)

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

- [go/client/README](go/client/README.md)

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
