# Install

- Protobuf
- Go
- Python

## Protobuf

- vscode plugins
  - [vscode-proto3](https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3)
  - [clang-format](https://marketplace.visualstudio.com/items?itemName=xaver.clang-format)

```bash
brew install clang-format
```

---

## Go

- Go
  - [gvm](https://github.com/moovweb/gvm)
- Protocol buffer 3+
  - [install](https://grpc.io/docs/protoc-installation/)
  
### Install Protoc

#### Ubuntu

```bash
sudo apt install -y protobuf-compiler
```

#### MacPort

```bash
sudo port install protobuf3-cpp
```

#### HomeBrew

```bash
brew install protobuf
```

#### Manually

- [protocolbuffers/protobuf](https://github.com/protocolbuffers/protobuf)
   - [release](https://github.com/protocolbuffers/protobuf/releases/latest)

Download pre-built binary: `protoc-XX.X-osx-x86_64.zip`

```bash
sudo mv protoc-XX/ /usr/local
sudo cp -R protoc-XX/include /usr/local/include
sudo chmod u+x /usr/local/protoc-XX/bin/protoc
sudo xattr -d com.apple.quarantine /usr/local/protoc-XX/bin/protoc
```

Add `PATH` in `.zprofile`:

```bash
### Protobuf
export PATH="$PATH:/usr/local/protoc-XX/bin"
```

### Check Versions

```bash
go version # go version go1.18 darwin/amd64
protoc --version # libprotoc 3.21.2
```

### Go plugins for the protocol compiler

- [protocolbuffers/protobuf-go/cmd/protoc-gen-go](https://github.com/protocolbuffers/protobuf-go/tree/master/cmd/protoc-gen-go)
- [grpc/grpc-go/cmd/protoc-gen-go-grpc](https://github.com/grpc/grpc-go/tree/master/cmd/protoc-gen-go-grpc)

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```bash
export PATH="$PATH:$(go env GOPATH)/bin"

echo $PATH | grep -E "$(go env GOPATH)/bin"
```
---

## Python

- Python 3.5+
- pip 9.0.1+

```bash
pip install --upgrade pip
pip install grpcio
pip install grpcio-tools
```

