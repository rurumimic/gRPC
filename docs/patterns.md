# gRPC patterns

[Core concepts, architecture and lifecycle](https://grpc.io/docs/what-is-grpc/core-concepts/)

---

src: [Order Service](../src/orderservice/README.md)

---

- Unary RPC
- Server streaming RPC
- Client streaming RPC
- Bidirectional streaming RPC

## Unary RPC

- the client sends a **single request** to the server and gets a **single response** back
- just like a normal function call

```protobuf
rpc SayHello(HelloRequest) returns (HelloResponse);
```

## Server streaming RPC

- the client sends **a request** to the server and gets a **stream to read a sequence of messages** back
- the client reads from the returned stream until there are no more messages
- **gRPC guarantees message ordering** within an individual RPC call

```protobuf
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);
```
## Client streaming RPC

- the client writes **a sequence of messages** and sends them to the server, again using a provided stream
- the client waits for the server to read them and return its **response**
- **gRPC guarantees message ordering** within an individual RPC call

```protobuf
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);
```
## Bidirectional streaming RPC

- **both** sides send **a sequence of messages** using a **read-write stream**
- the **order of messages** in each stream **is preserved**

```protobuf
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse);
```