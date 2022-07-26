# Documentation

<img src="https://www.oreilly.com/library/view/grpc-up-and/9781492058328/assets/grpc_0101.png" style="max-width: 700px;">

---

- [gRPC](grpc.md)
  - docs: [microsoft](docs.microsoft.md)
  - docs: [aws](docs.aws.md)
- [Helloworld](../src/helloworld/README.md)
- [Product Info Service](../src/productinfo/README.md)
  - Create go modules: server, client
  - Write and compile a protobuf for go interface
  - Build and run server and client
- gRPC [patterns](patterns.md): [Order Service](../src/orderservice/README.md)
  - Unary RPC
  - Server streaming RPC
  - Client streaming RPC
  - Bidirectional streaming RPC
- [Basic](basic.md)
  - RPC
  - Message Encoding
  - Length-Prefix Message Framing
  - HTTP/2
  - gRPC Layer
- [Advanced](advanced.md)
  - [Interceptors](advanced.md#interceptors)
    - server/client 
    - unary/stream interceptor
  - [Deadline](advanced.md#deadline)
  - [Cancellation](advanced.md#cancellation)
  - [Error Handling](advanced.md#error-handling)
  - [Metadata](advanced.md#metadata)
  - [Load Balancing](advanced.md#load-balancing)
  - [Multiplexing](advanced.md#multiplexing)
  - [Compression](advanced.md#compression)
- Secure
- Production
- Ecosystem
  - [microservices architecture](microservices.md)
