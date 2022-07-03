# productinfo

- [ProductInfo Service and Client](https://github.com/grpc-up-and-running/samples/tree/master/ch02)

## languages

- [go](go/README.md)
- [python](python/README.md)

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

---

## Mix

### Go Server + Python Client

#### Go Server

```bash
cd ./productinfo/go/server
bin/server
```

```bash
2022/07/03 23:59:48 Product f20d7937-6e30-450f-b239-590cd070b299 : Apple iPhone 11 - Added.
2022/07/03 23:59:48 Product f20d7937-6e30-450f-b239-590cd070b299 : Apple iPhone 11 - Retrieved.
```

#### Python Client

```bash
cd ./productinfo/python/client
python client.py
```

```bash
add product: response value: "f20d7937-6e30-450f-b239-590cd070b299"

get product: response id: "f20d7937-6e30-450f-b239-590cd070b299"
name: "Apple iPhone 11"
description: "Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode."
price: 699.0
```

### Python Server + Go Client

#### Python Server

```bash
cd ./productinfo/python/server
python server.py
```

```bash
Starting server. Listening on port 50051.
```

```bash
addProduct:request id: "4cc9c0a4-fae1-11ec-a6f0-8c85907c063a"
name: "Apple iPhone 11"
description: "Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode."
price: 699.0

addProduct:response value: "4cc9c0a4-fae1-11ec-a6f0-8c85907c063a"

getProduct:request value: "4cc9c0a4-fae1-11ec-a6f0-8c85907c063a"

getProduct:response id: "4cc9c0a4-fae1-11ec-a6f0-8c85907c063a"
name: "Apple iPhone 11"
description: "Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode."
price: 699.0
```

#### Go Client

```bash
cd ./productinfo/go/client
bin/client
```

```bash
2022/07/04 00:03:30 Product ID: 4cc9c0a4-fae1-11ec-a6f0-8c85907c063a added successfully
2022/07/04 00:03:30 Product: id:"4cc9c0a4-fae1-11ec-a6f0-8c85907c063a" name:"Apple iPhone 11" description:"Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode." price:699
```
