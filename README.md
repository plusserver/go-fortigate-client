# go-fortigate-client

Fortigate GO API client.

How to use:

```go
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
			{Ip: "10.90.250.1", Port: 80},
			{Ip: "10.90.250.2", Port: 80},
			{Ip: "10.90.250.3", Port: 80},
			{Ip: "10.90.250.4", Port: 80},
		},
	}
	
err := c.CreateVIP(vip)
	
...

err := c.DeleteVIP("myvip")

```

See `fgcmd.go` for more examples.
