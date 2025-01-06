# User-RPC-Service

## Installation

Compile .proto file by folder_name.

```bash
protoc --proto_path=./proto --go_out=. --go-grpc_out=. ./proto/{folder_name}/*.proto
```

## Usage
```bash
make run
```