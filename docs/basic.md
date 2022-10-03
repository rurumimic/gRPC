# Basic

- gRPC: Up and Running by Kasun Indrasiri, Danesh Kuruppu
  - [Chapter 4. gRPC: Under the Hood](https://www.oreilly.com/library/view/grpc-up-and/9781492058328/ch04.html)
- Protobuf
  - [Encoding](https://developers.google.com/protocol-buffers/docs/encoding)
- Learning HTTP/2 by Stephen Ludin, Javier Garza
  - [Appendix A. HTTP/2 Frames](https://www.oreilly.com/library/view/learning-http2/9781491962435/app01.html)
- [Protocol HTTP/2](https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md)
- [Existing Transports](https://grpc.github.io/grpc/core/md_doc_core_transport_explainer.html)

## RPC

![](/images/grpc_rpc.png)

1. Client Process: call a function in the generated stub
2. Client Stub: create an HTTP POST request
   - encoded message
   - content-type prefix: `application/grpc`
3. HTTP request message → server
4. Server: examine message headers
   - hands over the message to the service stub
5. Service Stub: parse the message bytes
   - into language-specific data structures
6. Service: call a local function
7. Server: send back to the client
   - encoded response message from the service function

---

## Message Encoding

### Encoded Message

<img src="/images/protobuf_message.png" width="600px">

#### Wire Types

| Type | Meaning          | Used For                                                 |
| ---- | ---------------- | -------------------------------------------------------- |
| 0    | Varint           | int32, int64, uint32, uint64, sint32, sint64, bool, enum |
| 1    | 64-bit           | fixed64, sfixed64, double                                |
| 2    | Length-delimited | string, bytes, embedded messages, packed repeated fields |
| 3    | Start            | group	groups (deprecated)                                |
| 4    | End              | group	groups (deprecated)                                |
| 5    | 32-bit           | fixed32, sfixed32, float                                 |

#### Tag Value

<img src="/images/protobuf_message_field.png" width="300px">

```go
Tag Value = (field_number << 3) | wire_type
```

#### Simple Message: varint

```protobuf
message Test1 {
  optional int32 a = 1;
}
```

```go
a := 150
```

```protobuf
08 96 01
```

dropping the MSB:

```protobuf
08 = 000 1000
  →  0001 000
  → field wire
```

- field number: 1
- wire type: 1

```protobuf
96 01 = 1001 0110     0000 0001
       → 000 0001  ++  001 0110 (drop the MSB and reverse the groups of 7 bits)
       → 10010110
       → 128 + 16 + 4 + 2 = 150
```

- value: 150

#### String Message

```protobuf
message Test2 {
  optional string b = 2;
}
```

```go
b := "testing"
```

```protobuf
12 07 [74 65 73 74 69 6e 67]
```

```protobuf
0x12
→ 0001 0010  (binary representation)
→ 00010 010  (regroup bits)
→ field_number = 2, wire_type = 2
```

---

## Length-Prefix Message Framing

<img src="/images/length_fixed_frame.png" width="600px">

- **Compressed Flag**: 1 byte
  - 0: no encoding
  - 1: compressed using the mechanism declared in the Message-Encoding header. declared in HTTP transport.
- **Message size**: 4 bytes
  - 4 bytes = 32 bits = 2^32 = 4,294,967,296 = 4 GB
- **Message Binary**
  - Big-endian

### transport implementations

- HTTP/2
- [Cronet](https://github.com/grpc/grpc/tree/master/src/core/ext/transport/cronet)
- [in-process](https://github.com/grpc/grpc/tree/master/src/core/ext/transport/inproc)

---

## HTTP/2

doc: [Protocol HTTP/2](https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md)

### HTTP/2 Frame Header

![](/images/http2_header.png)


### Request Message

- Request Headers
- Length-Prefixed Message
- End of Stream Flag

```text
HEADERS (flags = END_HEADERS)
:method = POST
:scheme = http
:path = /ProductInfo/getProduct
:authority = abc.com
te = trailers
grpc-timeout = 1S
content-type = application/grpc
grpc-encoding = gzip
authorization = Bearer xxxxxx
```

doc: [requests](https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests)

- method: always POST
- scheme: http / https
- path: /**ServiceName**/**MethodName**
- authority: virtual host name
- [te](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/TE): trailers = grpc
- `application/grpc`: [415 Unsupported Media Type](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/415)
- encoding: identity / gzip / deflate / snappy / {custom}


### Response Message

- Response Headers
- [Length-Prefixed Message]
- Trailers

```text
HEADERS (flags = END_HEADERS)
:status = 200
grpc-encoding = gzip
content-type = application/grpc
```

Trailer:

```text
HEADERS (flags = END_STREAM, END_HEADERS)
grpc-status = 0
grpc-message = xxxxxx
```

### Patterns

#### Unary RPC

<img src="/images/grpc_http2_unary.png" width="900px">

- End of Stream → half-close the connection

#### Server Streaming RPC

<img src="/images/grpc_http2_server.png" width="900px">

#### Client Streaming RPC

<img src="/images/grpc_http2_client.png" width="900px">

#### Bidirectional Streaming RPC

<img src="/images/grpc_http2_bid.png" width="900px">

---

## Architecture

<img src="/images/grpc_core.png" width="600px">

- Application Layer
  - Application Logic
    - Call Data Encoding Source Code
  - Data Encoding Logic
    - Protobuf Compiler → Data Encoding Source Code
- Framework Layer
  - provides extensions
    - authentication filter
    - deadline filter
