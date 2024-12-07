# Proxy

## 1. Run Python gRPC server

### 1.1. Install Python dependencies

```bash
cd python/echo
uv sync
source .venv/bin/activate
```

### 1.2. Run Python gRPC server

```bash
(.venv) python -m echo

Starting server. Listening on port 50051.
```

## 2. Run Go REST API server

```bash
cd go/server
```

```bash
go run .

gRPC client is connected to the server
Server is running on port 3000
```

## 3. Send a message

### 3.1. Check REST API server

client:

```bash
curl -XGET http://localhost:3000

Hello World!
```

go api server:

```bash
Server is running on port 3000
2024/12/07 20:40:26 "GET http://localhost:3000/ HTTP/1.1" from 127.0.0.1:54070 - 200 12B in 21.406µs
```

### 3.2. Send a message to REST API server

client:

```bash
curl -XPOST http://localhost:3000/echo -d '{"title":"Hello?","content":"Echo my message!"}'

Hello?
```

go api server:

```bash
Title:  Hello?
Content:  Echo my message!
2024/12/07 20:40:49 Echo: Hello?
2024/12/07 20:40:49 "POST http://localhost:3000/echo HTTP/1.1" from 127.0.0.1:45346 - 200 6B in 2.041809ms
```

python grpc server:

```bash
sendMessage:request title: "Hello?"
content: "Echo my message!"
sendMessage:response title: "Hello?"
```

