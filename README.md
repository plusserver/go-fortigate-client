# go-fortigate-client

Fortigate GO API client. Supports operations on "firewall", "certificate" and "vpn" for now.

How to use:

```go
import "github.com/Nexinto/go-fortigate-client/fortigate"

c := fortigate.NewWebClient(fortigate.WebClient{
	URL: os.Getenv("FORTIGATE_URL"),
	ApiKey: os.Getenv("FORTIGATE_API_KEY")})

vip := &fortigate.VIP{
		Name:            "myvip",
		Type:            fortigate.VIPTypeServerLoadBalance,
		LdbMethod:       fortigate.VIPLdbMethodRoundRobin,
		PortmappingType: fortigate.VIPPortmappingType1To1,
		Extintf:         "any",
		ServerType:      fortigate.VIPServerTypeTcp,
		Comment:         "my service",
		Extip:           "10.90.250.1",
		Extport:         "80",
		Realservers: []fortigate.VIPRealservers{
			{Ip: "10.90.251.1", Port: 80},
			{Ip: "10.90.251.2", Port: 80},
			{Ip: "10.90.251.3", Port: 80},
			{Ip: "10.90.251.4", Port: 80},
		},
	}
	
id, err := c.CreateVIP(vip)

...

err = c.UpdateVIP(vip)

...

err := c.DeleteVIP(id)

```

See `fgcmd.go` for more examples.

Create a Fake client for testing:

```go

c := fortigate.NewFakeClient()

```