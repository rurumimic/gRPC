# order service

- [gRPC Communication Patterns](https://github.com/grpc-up-and-running/samples/tree/master/ch03)

---

## Get start

```bash
orderservice
├── go
│   ├── client
│   └── server
└── proto
    └── order_management.proto
```

server:

```bash
cd go/server
go mod init ordermgt/server
```

client:

```bash
cd go/client
go mod init ordermgt/client
```

---

## Compile Proto

### Server

in `./orderservice`:

```bash
protoc -I proto \
  proto/order_management.proto \
  --go_out=go/server \
  --go-grpc_out=go/server
```

### Client

in `./orderservice`:

```bash
protoc -I proto \
  proto/order_management.proto \
  --go_out=go/client \
  --go-grpc_out=go/client
```

### Output

```bash
orderservice
├── go
│   ├── client
│   │   ├── ecommerce
│   │   │   ├── order_management.pb.go
│   │   │   └── order_management_grpc.pb.go
│   │   └── go.mod
│   └── server
│       ├── ecommerce
│       │   ├── order_management.pb.go
│       │   └── order_management_grpc.pb.go
│       └── go.mod
└── proto
    └── order_management.proto
```

---

## Programs

- [go/server/main.go](server/main.go)
- [go/client/main.go](client/main.go)

in `go/server` and `go/client`:

```bash
go mod tidy
```

---

## Run

in `go/server` and `go/client`:

```bash
go run main.go
```

--

## Build

in `go/server` and `go/client`:

```bash
go build -v -o bin/main
bin/main
```

---

## Output

### Server

#### AddOrder

```bash
2022/07/09 20:45:47 Order Added. ID : 101
```

#### SearchOrders 'Google'

```bash
2022/07/09 20:45:47 102{{{} [] [] <nil>} 0 [] 102 [Google Pixel 3A Mac Book Pro]  1800 Mountain View, CA}
2022/07/09 20:45:47 Google Pixel 3A
2022/07/09 20:45:47 Matching Order Found : 102
2022/07/09 20:45:47 103{{{} [] [] <nil>} 0 [] 103 [Apple Watch S4]  400 San Jose, CA}
2022/07/09 20:45:47 Apple Watch S4
2022/07/09 20:45:47 104{{{} [] [] <nil>} 0 [] 104 [Google Home Mini Google Nest Hub]  400 Mountain View, CA}
2022/07/09 20:45:47 Google Home Mini
2022/07/09 20:45:47 Matching Order Found : 104
2022/07/09 20:45:47 105{{{} [] [] <nil>} 0 [] 105 [Amazon Echo]  30 San Jose, CA}
2022/07/09 20:45:47 Amazon Echo
2022/07/09 20:45:47 106{{{} [] [] <nil>} 0 [] 106 [Amazon Echo Apple iPhone XS]  300 Mountain View, CA}
2022/07/09 20:45:47 Amazon Echo
2022/07/09 20:45:47 Apple iPhone XS
2022/07/09 20:45:47 101{{{} [] [] 0xc0001f38c0} 0 [] 101 [iPhone XS Mac Book Pro]  2300 San Jose, CA}
2022/07/09 20:45:47 iPhone XS
2022/07/09 20:45:47 Mac Book Pro
```

#### UpdateOrders

```bash
2022/07/09 20:45:47 Order ID : 102 - Updated
2022/07/09 20:45:47 Order ID : 103 - Updated
2022/07/09 20:45:47 Order ID : 104 - Updated
```

#### ProcessOrders

```bash
2022/07/09 20:45:47 Reading Proc order : value:"102"
2022/07/09 20:45:47 1cmb - Mountain View, CA
2022/07/09 20:45:47 Reading Proc order : value:"103"
2022/07/09 20:45:47 1cmb - San Jose, CA
2022/07/09 20:45:47 Reading Proc order : value:"104"
2022/07/09 20:45:47 Shipping : cmb - Mountain View, CA -> 2
2022/07/09 20:45:47 Shipping : cmb - San Jose, CA -> 1
2022/07/09 20:45:48 Reading Proc order : value:"101"
2022/07/09 20:45:48 1cmb - San Jose, CA
2022/07/09 20:45:48 Reading Proc order : <nil>
2022/07/09 20:45:48 EOF : <nil>
```

### Client

#### AddOrder 101

```bash
2022/07/09 20:45:36 AddOrder Response -> Order Added: 101
```

#### GetOrder 106

```bash
2022/07/09 20:45:36 GetOrder Response -> : id:"106" items:"Amazon Echo" items:"Apple iPhone XS" price:300 destination:"Mountain View, CA"
```

#### SearchOrders 'Google'

```bash
2022/07/09 20:45:36 Search Result : id:"102" items:"Google Pixel 3A" items:"Mac Book Pro" price:1800 destination:"Mountain View, CA"
2022/07/09 20:45:36 Search Result : id:"104" items:"Google Home Mini" items:"Google Nest Hub" price:400 destination:"Mountain View, CA"
2022/07/09 20:45:36 EOF
```

#### UpdateOrders

```bash
2022/07/09 20:45:36 Update Orders Res : value:"Orders processed Updated Order IDs : 102, 103, 104, "
```

#### ProcessOrders

```bash
2022/07/09 20:45:36 Combined shipment : %!(EXTRA []*ecommerce.Order=[id:"102" items:"Google Pixel 3A" items:"Google Pixel Book" price:1100 destination:"Mountain View, CA" id:"104" items:"Google Home Mini" items:"Google Nest Hub" items:"iPad Mini" price:2200 destination:"Mountain View, CA"])
2022/07/09 20:45:36 Combined shipment : %!(EXTRA []*ecommerce.Order=[id:"103" items:"Apple Watch S4" items:"Mac Book Pro" items:"iPad Pro" price:2800 destination:"San Jose, CA"])
2022/07/09 20:45:37 Combined shipment : %!(EXTRA []*ecommerce.Order=[id:"101" items:"iPhone XS" items:"Mac Book Pro" price:2300 destination:"San Jose, CA"])
```
