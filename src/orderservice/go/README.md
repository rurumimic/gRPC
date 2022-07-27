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

#### AddOrder: Unary

```bash
2022/07/09 20:45:47 Order Added. ID : 101
```

#### GetOrder: Unary

pass

#### SearchOrders 'Google': Server Streaming

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

#### UpdateOrders: Client Streaming

```bash
2022/07/09 20:45:47 Order ID : 102 - Updated
2022/07/09 20:45:47 Order ID : 103 - Updated
2022/07/09 20:45:47 Order ID : 104 - Updated
```

#### ProcessOrders: Bi directional Streaming

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

#### AddOrder 101: Unary

```bash
2022/07/09 20:45:36 AddOrder Response -> Order Added: 101
```

#### GetOrder 106: Unary

```bash
2022/07/09 20:45:36 GetOrder Response -> : id:"106" items:"Amazon Echo" items:"Apple iPhone XS" price:300 destination:"Mountain View, CA"
```

#### SearchOrders 'Google': Server Streaming

```bash
2022/07/09 20:45:36 Search Result : id:"102" items:"Google Pixel 3A" items:"Mac Book Pro" price:1800 destination:"Mountain View, CA"
2022/07/09 20:45:36 Search Result : id:"104" items:"Google Home Mini" items:"Google Nest Hub" price:400 destination:"Mountain View, CA"
2022/07/09 20:45:36 EOF
```

#### UpdateOrders: Client Streaming

```bash
2022/07/09 20:45:36 Update Orders Res : value:"Orders processed Updated Order IDs : 102, 103, 104, "
```

#### ProcessOrders: Bidrectional Streaming

```bash
2022/07/09 20:45:36 Combined shipment : %!(EXTRA []*ecommerce.Order=[id:"102" items:"Google Pixel 3A" items:"Google Pixel Book" price:1100 destination:"Mountain View, CA" id:"104" items:"Google Home Mini" items:"Google Nest Hub" items:"iPad Mini" price:2200 destination:"Mountain View, CA"])
2022/07/09 20:45:36 Combined shipment : %!(EXTRA []*ecommerce.Order=[id:"103" items:"Apple Watch S4" items:"Mac Book Pro" items:"iPad Pro" price:2800 destination:"San Jose, CA"])
2022/07/09 20:45:37 Combined shipment : %!(EXTRA []*ecommerce.Order=[id:"101" items:"iPhone XS" items:"Mac Book Pro" price:2300 destination:"San Jose, CA"])
```

---

## Advanced

## Interceptor

### Server

#### AddOrder: Unary

```bash
2022/07/23 14:56:26 ======= [Server Interceptor]  /ecommerce.OrderManagement/addOrder
2022/07/23 14:56:26  Pre Proc Message : id:"101" items:"iPhone XS" items:"Mac Book Pro" price:2300 destination:"San Jose, CA"
2022/07/23 14:56:26 Order Added. ID : 101
2022/07/23 14:56:26  Post Proc Message : id:"101" items:"iPhone XS" items:"Mac Book Pro" price:2300 destination:"San Jose, CA"
```

#### GetOrder: Unary

```bash
2022/07/23 14:56:26 ======= [Server Interceptor]  /ecommerce.OrderManagement/getOrder
2022/07/23 14:56:26  Pre Proc Message : value:"106"
2022/07/23 14:56:26  Post Proc Message : value:"106"
```

#### SearchOrders 'Google': Server Streaming

```bash
# Stream
# func orderServerStreamInterceptor
2022/07/23 15:25:40 ====== [Server Stream Interceptor]  /ecommerce.OrderManagement/searchOrders

# func RecvMsg
2022/07/23 15:25:40 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *wrapperspb.StringValue) at 2022-07-23T15:25:40+09:00

# func SearchOrders
2022/07/23 15:25:40 103{{{} [] [] <nil>} 0 [] 103 [Apple Watch S4]  400 San Jose, CA}
2022/07/23 15:25:40 Apple Watch S4
2022/07/23 15:25:40 104{{{} [] [] <nil>} 0 [] 104 [Google Home Mini Google Nest Hub]  400 Mountain View, CA}
2022/07/23 15:25:40 Google Home Mini

# func SendMsg
2022/07/23 15:25:40 ====== [Server Stream Interceptor Wrapper] Send a message (Type: *ecommerce.Order) at 2022-07-23T15:25:40+09:00
2022/07/23 15:25:40 Matching Order Found : 104

# func SearchOrders
2022/07/23 15:25:40 105{{{} [] [] <nil>} 0 [] 105 [Amazon Echo]  30 San Jose, CA}
2022/07/23 15:25:40 Amazon Echo
2022/07/23 15:25:40 106{{{} [] [] <nil>} 0 [] 106 [Amazon Echo Apple iPhone XS]  300 Mountain View, CA}
2022/07/23 15:25:40 Amazon Echo
2022/07/23 15:25:40 Apple iPhone XS
2022/07/23 15:25:40 101{{{} [] [] 0xc0000fb8c0} 0 [] 101 [iPhone XS Mac Book Pro]  2300 San Jose, CA}
2022/07/23 15:25:40 iPhone XS
2022/07/23 15:25:40 Mac Book Pro
2022/07/23 15:25:40 102{{{} [] [] <nil>} 0 [] 102 [Google Pixel 3A Mac Book Pro]  1800 Mountain View, CA}
2022/07/23 15:25:40 Google Pixel 3A

# func SendMsg
2022/07/23 15:25:40 ====== [Server Stream Interceptor Wrapper] Send a message (Type: *ecommerce.Order) at 2022-07-23T15:25:40+09:00
2022/07/23 15:25:40 Matching Order Found : 102
```

#### UpdateOrders: Client Streaming

```bash
# Stream
2022/07/23 15:34:20 ====== [Server Stream Interceptor]  /ecommerce.OrderManagement/updateOrders

# Recv
2022/07/23 15:34:20 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *ecommerce.Order) at 2022-07-23T15:34:20+09:00
2022/07/23 15:34:20 Order ID : 102 - Updated

2022/07/23 15:34:20 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *ecommerce.Order) at 2022-07-23T15:34:20+09:00
2022/07/23 15:34:20 Order ID : 103 - Updated

2022/07/23 15:34:20 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *ecommerce.Order) at 2022-07-23T15:34:20+09:00
2022/07/23 15:34:20 Order ID : 104 - Updated

# Recv: io.EOF
2022/07/23 15:34:20 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *ecommerce.Order) at 2022-07-23T15:34:20+09:00

# return stream.SendAndClose
2022/07/23 15:34:20 ====== [Server Stream Interceptor Wrapper] Send a message (Type: *wrapperspb.StringValue) at 2022-07-23T15:34:20+09:00
```

#### ProcessOrders: Bi directional Streaming

```bash
2022/07/23 15:39:55 ====== [Server Stream Interceptor]  /ecommerce.OrderManagement/processOrders

2022/07/23 15:39:55 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *wrapperspb.StringValue) at 2022-07-23T15:39:55+09:00
2022/07/23 15:39:55 Reading Proc order : value:"102"
2022/07/23 15:39:55 1cmb - Mountain View, CA

2022/07/23 15:39:55 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *wrapperspb.StringValue) at 2022-07-23T15:39:55+09:00
2022/07/23 15:39:55 Reading Proc order : value:"103"
2022/07/23 15:39:55 1cmb - San Jose, CA

2022/07/23 15:39:55 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *wrapperspb.StringValue) at 2022-07-23T15:39:55+09:00
2022/07/23 15:39:55 Reading Proc order : value:"104"

# stream.Send
2022/07/23 15:39:55 Shipping : cmb - Mountain View, CA -> 2
2022/07/23 15:39:55 ====== [Server Stream Interceptor Wrapper] Send a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T15:39:55+09:00

# stream.Send
2022/07/23 15:39:55 Shipping : cmb - San Jose, CA -> 1
2022/07/23 15:39:55 ====== [Server Stream Interceptor Wrapper] Send a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T15:39:55+09:00

2022/07/23 15:39:55 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *wrapperspb.StringValue) at 2022-07-23T15:39:55+09:00
2022/07/23 15:39:55 Reading Proc order : value:"101"
2022/07/23 15:39:55 1cmb - San Jose, CA

# Recv: io.EOF
2022/07/23 15:39:55 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *wrapperspb.StringValue) at 2022-07-23T15:39:55+09:00
2022/07/23 15:39:55 Reading Proc order : <nil>
2022/07/23 15:39:55 EOF : <nil>

# stream.Send
2022/07/23 15:39:55 ====== [Server Stream Interceptor Wrapper] Send a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T15:39:55+09:00
```

### Client

#### AddOrder 101: Unary

```bash
2022/07/23 16:18:54 Method : /ecommerce.OrderManagement/addOrder
2022/07/23 16:18:54 value:"Order Added: 101"
2022/07/23 16:18:54 AddOrder Response -> Order Added: 101
```

#### GetOrder 106: Unary

```bash
2022/07/23 16:18:54 Method : /ecommerce.OrderManagement/getOrder
2022/07/23 16:18:54 id:"106" items:"Amazon Echo" items:"Apple iPhone XS" price:300 destination:"Mountain View, CA"
2022/07/23 16:18:54 GetOrder Response -> : id:"106" items:"Amazon Echo" items:"Apple iPhone XS" price:300 destination:"Mountain View, CA"
```

#### SearchOrders 'Google': Server Streaming

```bash
# Stream
2022/07/23 18:46:48 ======= [Client Interceptor]  /ecommerce.OrderManagement/searchOrders

# Send: "Google"
2022/07/23 18:46:48 ====== [Client Stream Interceptor] Send a message (Type: *wrapperspb.StringValue) at 2022-07-23T18:46:48+09:00

# Recv
2022/07/23 18:46:48 ====== [Client Stream Interceptor] Receive a message (Type: *ecommerce.Order) at 2022-07-23T18:46:48+09:00
2022/07/23 18:46:48 Search Result : id:"102" items:"Google Pixel 3A" items:"Mac Book Pro" price:1800 destination:"Mountain View, CA"

# Recv
2022/07/23 18:46:48 ====== [Client Stream Interceptor] Receive a message (Type: *ecommerce.Order) at 2022-07-23T18:46:48+09:00
2022/07/23 18:46:48 Search Result : id:"104" items:"Google Home Mini" items:"Google Nest Hub" price:400 destination:"Mountain View, CA"

# Recv: io.EOF
2022/07/23 18:46:48 ====== [Client Stream Interceptor] Receive a message (Type: *ecommerce.Order) at 2022-07-23T18:46:48+09:00
2022/07/23 18:46:48 EOF
```

#### UpdateOrders: Client Streaming

```bash
# Stream
2022/07/23 18:49:54 ======= [Client Interceptor]  /ecommerce.OrderManagement/updateOrders

# Send
2022/07/23 18:49:54 ====== [Client Stream Interceptor] Send a message (Type: *ecommerce.Order) at 2022-07-23T18:49:54+09:00

# Send
2022/07/23 18:49:54 ====== [Client Stream Interceptor] Send a message (Type: *ecommerce.Order) at 2022-07-23T18:49:54+09:00

# Send
2022/07/23 18:49:54 ====== [Client Stream Interceptor] Send a message (Type: *ecommerce.Order) at 2022-07-23T18:49:54+09:00

# CloseAndRecv
2022/07/23 18:49:54 ====== [Client Stream Interceptor] Receive a message (Type: *wrapperspb.StringValue) at 2022-07-23T18:49:54+09:00
2022/07/23 18:49:54 Update Orders Res : value:"Orders processed Updated Order IDs : 102, 103, 104, "
```

#### ProcessOrders: Bidrectional Streaming

```bash
# Stream
2022/07/23 18:54:08 ======= [Client Interceptor]  /ecommerce.OrderManagement/processOrders

# Send
2022/07/23 18:54:08 ====== [Client Stream Interceptor] Send a message (Type: *wrapperspb.StringValue) at 2022-07-23T18:54:08+09:00

# Send
2022/07/23 18:54:08 ====== [Client Stream Interceptor] Send a message (Type: *wrapperspb.StringValue) at 2022-07-23T18:54:08+09:00

# Send
2022/07/23 18:54:08 ====== [Client Stream Interceptor] Send a message (Type: *wrapperspb.StringValue) at 2022-07-23T18:54:08+09:00

# Send
2022/07/23 18:54:08 ====== [Client Stream Interceptor] Send a message (Type: *wrapperspb.StringValue) at 2022-07-23T18:54:08+09:00

# Recv
2022/07/23 18:54:08 ====== [Client Stream Interceptor] Receive a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T18:54:08+09:00
2022/07/23 18:54:08 Combined shipment : [id:"103" items:"Apple Watch S4" items:"Mac Book Pro" items:"iPad Pro" price:2800 destination:"San Jose, CA"]

# Recv
2022/07/23 18:54:08 ====== [Client Stream Interceptor] Receive a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T18:54:08+09:00
2022/07/23 18:54:08 Combined shipment : [id:"102" items:"Google Pixel 3A" items:"Google Pixel Book" price:1100 destination:"Mountain View, CA" id:"104" items:"Google Home Mini" items:"Google Nest Hub" items:"iPad Mini" price:2200 destination:"Mountain View, CA"]

# Recv
2022/07/23 18:54:08 ====== [Client Stream Interceptor] Receive a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T18:54:08+09:00
2022/07/23 18:54:08 Combined shipment : [id:"101" items:"iPhone XS" items:"Mac Book Pro" price:2300 destination:"San Jose, CA"]

# Recv: io.EOF
2022/07/23 18:54:08 ====== [Client Stream Interceptor] Receive a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T18:54:08+09:00
```

---

## Deadlines

### Server

```bash
2022/07/23 19:33:19 ======= [Server Interceptor]  /ecommerce.OrderManagement/addOrder
2022/07/23 19:33:19  Pre Proc Message : id:"101"  items:"iPhone XS"  items:"Mac Book Pro"  price:2300  destination:"San Jose, CA"

# Sleep 5s
2022/07/23 19:33:19 Sleeping for:  5 s

# DEADLINE_EXCEEDED
2022/07/23 19:33:24 RPC has reached deadline exceeded state : context deadline exceeded

2022/07/23 19:33:24  Post Proc Message : id:"101"  items:"iPhone XS"  items:"Mac Book Pro"  price:2300  destination:"San Jose, CA"
```

### Client

```bash
# Deadline 2s
2022/07/23 19:33:19 Method : /ecommerce.OrderManagement/addOrder
2022/07/23 19:33:21

# DEADLINE_EXCEEDED
2022/07/23 19:33:21 Error Occured -> addOrder : DeadlineExceeded
```

---

## Cancellation

### Client side cancel

#### Server

```bash
2022/07/23 20:07:38 ====== [Server Stream Interceptor]  /ecommerce.OrderManagement/processOrders
2022/07/23 20:07:38 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *wrapperspb.StringValue) at 2022-07-23T20:07:38+09:00
2022/07/23 20:07:38 Reading Proc order : value:"102"
2022/07/23 20:07:38 1cmb - Mountain View, CA
2022/07/23 20:07:38  Context Cacelled for this stream: -> context canceled
2022/07/23 20:07:38 Stopped processing any more order of this stream!
2022/07/23 20:07:38 RPC failed with error context canceled
```

#### Client

```bash
2022/07/23 20:07:38 ======= [Client Interceptor]  /ecommerce.OrderManagement/processOrders
2022/07/23 20:07:38 ====== [Client Stream Interceptor] Send a message (Type: *wrapperspb.StringValue) at 2022-07-23T20:07:38+09:00
2022/07/23 20:07:38 ====== [Client Stream Interceptor] Send a message (Type: *wrapperspb.StringValue) at 2022-07-23T20:07:38+09:00
2022/07/23 20:07:38 ====== [Client Stream Interceptor] Send a message (Type: *wrapperspb.StringValue) at 2022-07-23T20:07:38+09:00
2022/07/23 20:07:38 RPC Status : context canceled
2022/07/23 20:07:38 ====== [Client Stream Interceptor] Send a message (Type: *wrapperspb.StringValue) at 2022-07-23T20:07:38+09:00
2022/07/23 20:07:38 ====== [Client Stream Interceptor] Receive a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T20:07:38+09:00
2022/07/23 20:07:38 Error Receiving messages rpc error: code = Canceled desc = context canceled
```

### Server side Cancel

#### Server

```bash
2022/07/23 20:10:10 ====== [Server Stream Interceptor]  /ecommerce.OrderManagement/processOrders
2022/07/23 20:10:10 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *wrapperspb.StringValue) at 2022-07-23T20:10:10+09:00
2022/07/23 20:10:10 Reading Proc order : value:"102"
2022/07/23 20:10:10 1cmb - Mountain View, CA
2022/07/23 20:10:10 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *wrapperspb.StringValue) at 2022-07-23T20:10:10+09:00
2022/07/23 20:10:10 Reading Proc order : value:"103"
2022/07/23 20:10:10 1cmb - San Jose, CA
2022/07/23 20:10:10 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *wrapperspb.StringValue) at 2022-07-23T20:10:10+09:00
2022/07/23 20:10:10 Reading Proc order : value:"104"
2022/07/23 20:10:10 Shipping : cmb - Mountain View, CA -> 2
2022/07/23 20:10:10 ====== [Server Stream Interceptor Wrapper] Send a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T20:10:10+09:00
2022/07/23 20:10:10 Shipping : cmb - San Jose, CA -> 1
2022/07/23 20:10:10 ====== [Server Stream Interceptor Wrapper] Send a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T20:10:10+09:00
2022/07/23 20:10:10 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *wrapperspb.StringValue) at 2022-07-23T20:10:10+09:00
2022/07/23 20:10:10 Reading Proc order : value:"101"
2022/07/23 20:10:10 1cmb - San Jose, CA
2022/07/23 20:10:10 ====== [Server Stream Interceptor Wrapper] Receive a message (Type: *wrapperspb.StringValue) at 2022-07-23T20:10:10+09:00
2022/07/23 20:10:10 Reading Proc order : <nil>
2022/07/23 20:10:10 EOF : <nil>
2022/07/23 20:10:10 ====== [Server Stream Interceptor Wrapper] Send a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T20:10:10+09:00
```

#### Client

```bash
2022/07/23 20:12:38 ======= [Client Interceptor]  /ecommerce.OrderManagement/processOrders
2022/07/23 20:12:38 ====== [Client Stream Interceptor] Send a message (Type: *wrapperspb.StringValue) at 2022-07-23T20:12:38+09:00
2022/07/23 20:12:38 ====== [Client Stream Interceptor] Send a message (Type: *wrapperspb.StringValue) at 2022-07-23T20:12:38+09:00
2022/07/23 20:12:38 ====== [Client Stream Interceptor] Send a message (Type: *wrapperspb.StringValue) at 2022-07-23T20:12:38+09:00
2022/07/23 20:12:38 ====== [Client Stream Interceptor] Send a message (Type: *wrapperspb.StringValue) at 2022-07-23T20:12:38+09:00
2022/07/23 20:12:38 ====== [Client Stream Interceptor] Receive a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T20:12:38+09:00
2022/07/23 20:12:38 Combined shipment : [id:"102" items:"Google Pixel 3A" items:"Google Pixel Book" price:1100 destination:"Mountain View, CA" id:"104" items:"Google Home Mini" items:"Google Nest Hub" items:"iPad Mini" price:2200 destination:"Mountain View, CA"]
2022/07/23 20:12:38 ====== [Client Stream Interceptor] Receive a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T20:12:38+09:00
2022/07/23 20:12:38 Combined shipment : [id:"103" items:"Apple Watch S4" items:"Mac Book Pro" items:"iPad Pro" price:2800 destination:"San Jose, CA"]
2022/07/23 20:12:38 ====== [Client Stream Interceptor] Receive a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T20:12:38+09:00
2022/07/23 20:12:38 Combined shipment : [id:"101" items:"iPhone XS" items:"Mac Book Pro" price:2300 destination:"San Jose, CA"]

# Server side cancel
2022/07/23 20:12:38 ====== [Client Stream Interceptor] Receive a message (Type: *ecommerce.CombinedShipment) at 2022-07-23T20:12:38+09:00
2022/07/23 20:12:38 Error Receiving messages EOF
```

---

## Error Handling

### Server

```bash
2022/07/26 19:51:11 ======= [Server Interceptor]  /ecommerce.OrderManagement/addOrder
2022/07/26 19:51:11  Pre Proc Message : id:"-1" items:"iPhone XS" items:"Mac Book Pro" price:2300 destination:"San Jose, CA"
2022/07/26 19:51:11 Order ID is invalid! -> Received Order ID -1
2022/07/26 19:51:11  Post Proc Message : id:"-1" items:"iPhone XS" items:"Mac Book Pro" price:2300 destination:"San Jose, CA"
```

### Client

```bash
2022/07/26 19:51:11 Method : /ecommerce.OrderManagement/addOrder
2022/07/26 19:51:11
2022/07/26 19:51:11 Invalid Argument Error : InvalidArgument
2022/07/26 19:51:11 Request Field Invalid: field:"ID" description:"Order ID received is not valid -1 : "
```

---

## Multiplexing

### Server

```bash
2022/07/27 19:48:20 ======= [Server Interceptor]  /helloworld.Greeter/SayHello
2022/07/27 19:48:20  Pre Proc Message : name:"gRPC Up and Running!"
2022/07/27 19:48:20 Greeter Service - SayHello RPC
2022/07/27 19:48:20  Post Proc Message : name:"gRPC Up and Running!"
# ...
```

### Client

```bash
2022/07/27 19:48:20 Method : /helloworld.Greeter/SayHello
2022/07/27 19:48:20 message:"Hello gRPC Up and Running!"
Greeting :  Hello gRPC Up and Running!
# ...
```

