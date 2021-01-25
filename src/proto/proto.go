//go:generate protoc --go_out=./grpc --go_opt=paths=source_relative --go_grpc_out=./grpc --go_grpc_opt=paths=source_relative ./proto.proto
package proto
