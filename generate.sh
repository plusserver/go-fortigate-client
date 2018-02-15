rm -f fortigate/{client,firewalladdress,firewallpolicy,vip}.go
go run fortigate/gen/generator.go
go fmt fortigate/*.go