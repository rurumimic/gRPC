# Advanced

- grpc-go: [features](https://github.com/grpc/grpc-go/tree/master/examples/features)

## Interceptors

- grpc: [Interceptors in gRPC-Web](https://grpc.io/blog/grpc-web-interceptor/)
- grpc-go: [intercetor.go](https://github.com/grpc/grpc-go/blob/master/interceptor.go)


![](https://grpc.io/img/grpc-web-interceptors.png)

- Unary Interceptor
- Stream Interceptor

### Server side interceptor

```go
grpc.NewServer(
  grpc.UnaryInterceptor(UnaryServerInterceptor),
  grpc.StreamInterceptor(ServerStreamInterceptor),
)
```

#### Unary Interceptor

`intercetor.go`:

```go
type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
```

server: [orderservice / server / main.go](/src/orderservice/go/server/main.go)

```go
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
  return handler(ctx, req)
}
```

#### Stream Interceptor

`intercetor.go`:

```go
type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error
```

type [grpc.ServerStream](https://pkg.go.dev/google.golang.org/grpc#ServerStream):

```go
type ServerStream interface {
	SetHeader(metadata.MD) error
	SendHeader(metadata.MD) error
	SetTrailer(metadata.MD)
	Context() context.Context
	SendMsg(m interface{}) error
	RecvMsg(m interface{}) error
}
```

server: [orderservice / server / main.go](/src/orderservice/go/server/main.go)

```go
type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	return w.ServerStream.SendMsg(m)
}
```

```go
func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}
```

```go
func ServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}
	return err
}
```

### Client side interceptor

```go
conn, err := grpc.Dial(
  address, 
  grpc.WithInsecure(),
  grpc.WithUnaryInterceptor(orderUnaryClientInterceptor),
  grpc.WithStreamInterceptor(clientStreamInterceptor)
)
```

#### Unary Interceptor

`intercetor.go`:

```go
type UnaryClientInterceptor func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error
```

client: [orderservice / client / main.go](/src/orderservice/go/client/main.go)

```go
func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return invoker(ctx, method, req, reply, cc, opts...)
}

```

#### Stream Interceptor

`intercetor.go`:

```go
type StreamClientInterceptor func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, streamer Streamer, opts ...CallOption) (ClientStream, error)
```

type [grpc.ClientStream](https://pkg.go.dev/google.golang.org/grpc#ClientStream)

```go
type ClientStream interface {
	Header() (metadata.MD, error)
	Trailer() metadata.MD
	CloseSend() error
	Context() context.Context
	SendMsg(m interface{}) error
	RecvMsg(m interface{}) error
}
```

client: [orderservice / client / main.go](/src/orderservice/go/client/main.go)

```go
type wrappedStream struct {
	grpc.ClientStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	return w.ClientStream.SendMsg(m)
}
```

```go
func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}
```

```go
func StreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(s), nil
}
```

---

## Deadline

- grpc: [gRPC and Deadlines](https://grpc.io/blog/deadlines/)

Error: `DEADLINE_EXCEEDED`

[Setting a deadline](https://grpc.io/blog/deadlines/#setting-a-deadline):

```go
clientDeadline := time.Now().Add(time.Duration(*deadlineMs) * time.Millisecond)
ctx, cancel := context.WithDeadline(ctx, clientDeadline)
```

[Checking deadlines](https://grpc.io/blog/deadlines/#checking-deadlines):

```go
if ctx.Err() == context.Canceled {
	return status.New(codes.Canceled, "Client cancelled, abandoning.")
}
```

[Adjusting deadlines](https://grpc.io/blog/deadlines/#adjusting-deadlines):

```go
var deadlineMs = flag.Int("deadline_ms", 20*1000, "Default deadline in milliseconds.")

ctx, cancel := context.WithTimeout(ctx, time.Duration(*deadlineMs) * time.Millisecond)
```

### Example

- server: [orderservice / server / main.go](/src/orderservice/go/server/main.go)
- client: [orderservice / client / main.go](/src/orderservice/go/client/main.go)

#### Server

```go
if ctx.Err() == context.DeadlineExceeded {
  log.Printf("RPC has reached deadline exceeded state : %s", ctx.Err())
  return nil, ctx.Err()
}
```

#### Client

```go
clientDeadline := time.Now().Add(time.Duration(2 * time.Second))
ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
defer cancel()
```

---

## Cancellation

### Example

#### Server

```go
if stream.Context().Err() == context.Canceled {
	log.Printf(" Context Cacelled for this stream: -> %s", stream.Context().Err())
	log.Printf("Stopped processing any more order of this stream!")
	return stream.Context().Err()
}
```

#### Client

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

```go
cancel()
log.Printf("RPC Status : %s", ctx.Err())
```

---

## Error Handling

- grpc: [Error handling](https://www.grpc.io/docs/guides/error/)
   - github: [statuscodes](https://github.com/grpc/grpc/blob/master/doc/statuscodes.md)
   - go: [codes](https://pkg.go.dev/google.golang.org/grpc/codes)

### Example

```go
import (
  epb "google.golang.org/genproto/googleapis/rpc/errdetails"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
)
```

#### Server

```go
if orderReq.Id == "-1" {
  errorStatus := status.New(codes.InvalidArgument, "Invalid information received")
  ds, err := errorStatus.WithDetails(
    &epb.BadRequest_FieldViolation{
      Field:       "ID",
      Description: fmt.Sprintf("Order ID received is not valid %s : %s", orderReq.Id, orderReq.Description),
    },
  )
  if err != nil {
    return nil, errorStatus.Err()
  }

  return nil, ds.Err()
}
```

#### Client

```go
res, invalidError := client.AddOrder(ctx, &invalidOrder)
if invalidError != nil {
  errorCode := status.Code(invliadError)
  if errorCode == codes.InvalidArgument {
    errorStatus := status.Convert(invlidError)
    for _, d := range errorStatus.Details() {
      switch info := d.(type) {
      case *epb.BadRequest_FieldViolation:
        log.Printf(info)
      default:
        log.Printf(info)
      }
    }
  }
}
```

---

## Multiplexing

protobuf: [helloworld](google.golang.org/grpc/examples/helloworld/helloworld)

```go
import (
  hello_pb "google.golang.org/grpc/examples/helloworld/helloworld"
)
```

```bash
go mod tidy
go run main.go
```

### Server

```go
pb.RegisterOrderManagementServer(s, &server{})
pb.RegisterGreeterServer(s, &helloServer{})
```

```go
type helloServer struct{}

func (s *helloServer) SayHello(ctx context.Context, in *hello_pb.HelloRequest) (*hello_pb.HelloReply, error) {
	log.Printf("Greeter Service - SayHello RPC")
	return &hello_pb.HelloReply{Message: "Hello " + in.Name}, nil
}
```

### Client

```go
helloClient := hello_pb.NewGreeterClient(conn)

/* HelloWorld */
helloCtx, helloCancel := context.WithTimeout(context.Background(), time.Second)
defer helloCancel()
helloResponse, err := helloClient.SayHello(helloCtx, &hello_pb.HelloRequest{Name: "gRPC Up and Running!"})
if err != nil {
  log.Fatalf("helloClient.SayHello(_) = _, %v", err)
}
fmt.Println("Greeting : ", helloResponse.Message)
```

---

## Metadata

- github: [metadata](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md)
- pkg: [metadata](https://pkg.go.dev/google.golang.org/grpc/metadata)

a map from string to a list of strings: `{'key': ['str', ...], }`

```go
type MD map[string][]string
```

### Example

- order service: [README.md](/src/orderservice/go/README.md#metadata)
- server: [orderservice / server / main.go](/src/orderservice/go/server/main.go)
- client: [orderservice / client / main.go](/src/orderservice/go/client/main.go)

```go
import "google.golang.org/grpc/metadata"
```

#### Server

unary rpc:

```go
func (s *server) AddOrder(ctx context.Context, orderReq *pb.Order) (*wrapper.StringValue, error) {
  // ...

	// Read metadata from client
	md, metadataAvailable := metadata.FromIncomingContext(ctx)
	if !metadataAvailable {
		return nil, status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	}
	if t, ok := md["timestamp"]; ok {
		fmt.Printf("timestamp from metadata:\n")
		for i, e := range t {
			fmt.Printf("====> Metadata %d. %s\n", i, e)
		}

		fmt.Println("Additional Metadata")
		for k, v := range md {
			fmt.Printf("====> Metadata %s. %s\n", k, v)
		}
	}

	// create and send a header
	header := metadata.New(map[string]string{"location": "San Jose", "timestamp": time.Now().Format(time.StampNano)})
	grpc.SendHeader(ctx, header)

  // ...
}
```

stream rpc:

```go
func (s *server) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {

	defer func() {
		trailer := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
		stream.SetTrailer(trailer)
	}()

	header := metadata.New(map[string]string{"location": "MTV", "timestamp": time.Now().Format(time.StampNano)})
	stream.SendHeader(header)

  // ...
}
```

#### Client

unary rpc:

```go
md := metadata.Pairs(
  "timestamp", time.Now().Format(time.StampNano),
  "kn", "vn",
)
ctx = metadata.NewOutgoingContext(ctx, md)
ctx = metadata.AppendToOutgoingContext(ctx, "k1", "v1", "k1", "v2", "k2", "v3")

var header, trailer metadata.MD

res, addErr := client.AddOrder(ctx, &order1, grpc.Header(&header), grpc.Trailer(&trailer))

if t, ok := header["timestamp"]; ok {
  log.Printf("timestamp from header:\n")
  for i, e := range t {
  	fmt.Printf(" %d. %s\n", i, e)
  }
} else {
  log.Fatal("timestamp expected but doesn't exist in header")
}
if l, ok := header["location"]; ok {
  log.Printf("location from header:\n")
  for i, e := range l {
  	fmt.Printf(" %d. %s\n", i, e)
  }
} else {
  log.Fatal("location expected but doesn't exist in header")
}
```

stream rpc:

```go
searchHeader, err := searchStream.Header()
if err == nil {
  for k, v := range searchHeader {
  	fmt.Printf("Header: %s. %s\n", k, v)
  }
}

searchTrailer := searchStream.Trailer()
for k, v := range searchTrailer {
  fmt.Printf("Trailer: %s. %s\n", k, v)
}
```

### Create a new metadata

#### New

- doc: [New](https://pkg.go.dev/google.golang.org/grpc/metadata#New)

```go
func New(m map[string]string) MD
```

```go
md := metadata.New(map[string]string{"key1": "val1", "key2": "val2"})
```

#### Paris

- doc: [Pairs](https://pkg.go.dev/google.golang.org/grpc/metadata#Pairs)

```go
func Pairs(kv ...string) MD
```

all the keys will be automatically converted to lowercase:

```go
md := metadata.Pairs(
    "key1", "val1",
    "key1", "val1-2", // "key1" will have map value []string{"val1", "val1-2"}
    "key2", "val2",
)
```

#### Binary data

this binary data will be encoded (base64) before sending,  
and will be decoded after being transferred.

simply add "-bin" suffix to the key:

```go
md := metadata.Pairs(
    "key", "string value",
    "key-bin", string([]byte{96, 102}),
)
```

### Retrieve metadata from context

```go
func (s *server) SomeRPC(ctx context.Context, in *pb.SomeRequest) (*pb.SomeResponse, err) {
    md, ok := metadata.FromIncomingContext(ctx)
    // do something with metadata
}
```

### Send and receive metadata - client side

- example: [client.go](https://github.com/grpc/grpc-go/blob/master/examples/features/metadata/client/main.go)

#### Send metadata

```go
md := metadata.Pairs(
  "timestamp", time.Now().Format(timestampFormat),
  "kn", "vn",
)

ctx := metadata.NewOutgoingContext(context.Background(), md)
ctx := metadata.AppendToOutgoingContext(ctx, "k1", "v1", "k1", "v2", "k2", "v3")

response, err := client.SomeRPC(ctx, someRequest) // Unary RPC
stream, err := client.SomeStreamingRPC(ctx) // Streaming RPC
```

or:

```go
md := metadata.Pairs("k1", "v1", "k1", "v2", "k2", "v3")
ctx := metadata.NewOutgoingContext(context.Background(), md)

// in an intererceptor
send, _ := metadata.FromOutgoingContext(ctx)
newMD := metadata.Pairs("k3", "v3")
ctx = metadata.NewOutgoingContext(ctx, metadata.Join(send, newMD))

response, err := client.SomeRPC(ctx, someRequest) // Unary RPC
stream, err := client.SomeStreamingRPC(ctx) // Streaming RPC
```

#### Recieve metadata

- func [Header](https://pkg.go.dev/google.golang.org/grpc#Header)
- func [Trailer](https://pkg.go.dev/google.golang.org/grpc#Trailer)

```go
var header, trailer metadata.MD

r, err := client.SomeRPC( // Unary RPC
  ctx,
  someRequest,
  grpc.Header(&header),
  grpc.Trailer(&trailer),
)

stream, err := client.SomeStreamingRPC(ctx) // Streaming RPC
header, err := stream.Header()
trailer, err := stream.Trailer()
```

### Send and receive metadata - server side

- example: [server.go](https://github.com/grpc/grpc-go/blob/master/examples/features/metadata/server/main.go)

#### Recieve metadata

```go
func (s *server) SomeRPC(ctx context.Context, in *pb.someRequest) (*pb.someResponse, error) {
    md, ok := metadata.FromIncomingContext(ctx)
    // do something with metadata
}

func (s *server) SomeStreamingRPC(stream pb.Service_SomeStreamingRPCServer) error {
    md, ok := metadata.FromIncomingContext(stream.Context()) // get context from stream
    // do something with metadata
}
```

#### Send metadata

```go
func (s *server) SomeRPC(ctx context.Context, in *pb.someRequest) (*pb.someResponse, error) {
    // create and send header
    header := metadata.Pairs("header-key", "val")
    grpc.SendHeader(ctx, header)
    // create and set trailer
    trailer := metadata.Pairs("trailer-key", "val")
    grpc.SetTrailer(ctx, trailer)
}

func (s *server) SomeStreamingRPC(stream pb.Service_SomeStreamingRPCServer) error {
    // create and send header
    header := metadata.Pairs("header-key", "val")
    stream.SendHeader(header)
    // create and set trailer
    trailer := metadata.Pairs("trailer-key", "val")
    stream.SetTrailer(trailer)
}
```

---

## Load Balancing

- [load balancer](/src/loadbalancer/README.md)

### Ref

- gRPC: [Name Resolution](https://grpc.github.io/grpc/core/md_doc_naming.html)
- examples
  - [name resolving](https://github.com/grpc/grpc-go/tree/master/examples/features/name_resolving)
    - [client.go](https://github.com/grpc/grpc-go/blob/master/examples/features/name_resolving/client/main.go)
  - [load balancing](https://github.com/grpc/grpc-go/blob/master/examples/features/load_balancing)
    - [server.go](https://github.com/grpc/grpc-go/blob/master/examples/features/load_balancing/server/main.go)
    - [client.go](https://github.com/grpc/grpc-go/blob/master/examples/features/load_balancing/client/main.go)
- doc
  - [Resolver](https://godoc.org/google.golang.org/grpc/resolver#Resolver)
  - [Builder](https://godoc.org/google.golang.org/grpc/resolver#Builder)

### Name Resolver

```go
const (
  exampleScheme = "example"
  exampleServiceName = "resolver.example.grpc.io"
)

var addrs = []string{"localhost:50051", "localhost:50052"}

type exampleResolverBuilder struct{}

func (*exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
  r := &exampleResolver {
    target: target,
    cc: cc,
    addrsStore: map[string][]string{
      exampleServiceName: addrs,
    },
  }

  r.start()

  return r, nil
}

func (*exampleResolverBuilder) Scheme() string { return exampleScheme }

type exampleResolver struct {
  target     resolver.Target
  cc         resolver.ClientConn
  addrsStore map[string][]string
}

func (r *exampleResolver) start() {
  addrStrs := r.addrsStore[r.target.Endpoint]
  addrs := make([]resolver.Address, len(addrStrs))
  for i, s := range addrStrs {
    addrs[i] = resolver.Address{Addr: s}
  }
  r.cc.UpdateState(resolver.State{Addresses: addrs})
}

func (*exampleResolver) ResolveNow(o resolver.ResolveNowOption) {}
func (*exampleResolver) Close() {}

func init() {
  resolver.Register(&exampleResolverBuilder{})
}
```

### Client Side Load Balancer

- Thick(Fat) Client
- Lookaside Load Balancer

`scheme:///server`

#### Client

pick first:

```go
// "pick_first" is the default, so there's no need to set the load balancing policy.
pickfirstConn, err := grpc.Dial(
  fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName),
  grpc.WithInsecure(),
)
if err != nil {
  log.Fatalf("did not connect: %v", err)
}
defer pickfirstConn.Close()

fmt.Println("--- calling helloworld.Greeter/SayHello with pick_first ---")
makeRPCs(pickfirstConn, 10)
```

round robin:

```go
// Make another ClientConn with round_robin policy.
roundrobinConn, err := grpc.Dial(
  fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName),
  grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`), // This sets the initial balancing policy.
  grpc.WithInsecure(),
)
if err != nil {
log.Fatalf("did not connect: %v", err)
}
defer roundrobinConn.Close()

fmt.Println("--- calling helloworld.Greeter/SayHello with round_robin ---")
makeRPCs(roundrobinConn, 10)
```

---

## Compression

- example: [compression](https://github.com/grpc/grpc-go/tree/master/examples/features/compression)
  - [server.go](https://github.com/grpc/grpc-go/blob/master/examples/features/compression/server/main.go)
  - [client.go](https://github.com/grpc/grpc-go/blob/master/examples/features/compression/client/main.go)

### Server

```go
import (
  _ "google.golang.org/grpc/encoding/gzip" // Install the gzip compressor
)
```

### Client

```go
import (
  "google.golang.org/grpc/encoding/gzip" // Install the gzip compressor
)

res, _ := client.AddOrder(ctx, &order1, grpc.UseCompressor(gzip.Name))
```
