# gRPC Demo in Go

Simple gRPC demonstration implementing all four communication patterns in Go.

## Prerequisites

- Go 1.23.4+
- Protocol Buffers compiler
- gRPC tools

Install required tools:
```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Project Structure
.
├── client/
│   └── main.go
├── proto/
│   ├── greet.proto
│   ├── greet.pb.go
│   └── greet_grpc.pb.go
├── server/
│   └── main.go
└── README.md
