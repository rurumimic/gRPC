# Load Balancer

## Code

- [server](go/server/main.go)
- [client](go/client/main.go)

### Server

```bash
cd go/server
go mod init lb/server
go mod tidy
```

### Client

```bash
cd go/server
go mod init lb/client
go mod tidy
```

---

## Output

```bash
go run main.go
```

### Server

```bash
2022/08/06 17:16:37 serving on :50052
2022/08/06 17:16:37 serving on :50051
```

### Client

```bash
--- calling helloworld.Greeter/SayHello with pick_first ---
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
--- calling helloworld.Greeter/SayHello with round_robin ---
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
```
