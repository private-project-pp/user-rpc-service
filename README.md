# User-RPC-Service

## Installation

Compile .proto file by folder_name.

```bash
protoc --proto_path=./proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/{folder_name}/*.proto
```

## Usage
```bash
make run
```