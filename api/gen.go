//go:generate protoc --proto_path=. --go_out=paths=source_relative:. --go-drpc_out=paths=source_relative:. foo.proto
package api
