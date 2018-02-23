rm fortigate/types.go
go run fortigate/gen/generator.go
go fmt fortigate/types.go
