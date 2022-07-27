# Advanced

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


---

## Error Handling


---

## Load Balancing


---

## Compression

