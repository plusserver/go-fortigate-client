package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Nexinto/go-fortigate-client/fortigate"
	"github.com/docopt/docopt-go"
)

func main() {

	usage := `fgcmd - fortigate command line utility

Usage:
  fgcmd vip list
  fgcmd vip show <name>
  fgcmd vip delete <name>
  fgcmd vip create <name> <ip>:<port> <realservers>
`

	c := fortigate.NewWebClient(fortigate.WebClient{URL: os.Getenv("FORTIGATE_URL"), ApiKey: os.Getenv("FORTIGATE_API_KEY")})

	opts, err := docopt.ParseDoc(usage)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", opts)

	if b, _ := opts.Bool("vip"); b {

		if b, _ := opts.Bool("list"); b {

			vips, err := c.ListVIPs()

			if err != nil {
				log.Fatalf("%s", err.Error())
			}

			for _, vip := range vips {
				fmt.Printf("%-30s %s:%s\n", vip.Name, vip.Extip, vip.Extport)
			}

		} else if b, _ := opts.Bool("show"); b {

			name, _ := opts.String("<name>")
			vip, err := c.GetVIP(name)
			if err != nil {
				log.Fatalf("%s", err.Error())
			}

			fmt.Printf("VIP Name: %s\nIP:       %s\nPort(s):  %s\n\nRealservers:\n", vip.Name, vip.Extip, vip.Extport)
			for _, rs := range vip.Realservers {
				fmt.Printf("%s:%d\n", rs.Ip, rs.Port)
			}

		} else if b, _ := opts.Bool("delete"); b {

			name, _ := opts.String("<name>")
			err := c.DeleteVIP(name)
			if err != nil {
				log.Fatalf("%s", err.Error())
			}

		} else if b, _ := opts.Bool("create"); b {

			name, _ := opts.String("<name>")
			vipp, _ := opts.String("<ip>:<port>")
			parts := strings.Split(vipp, ":")
			extip, extport := parts[0], parts[1]
			realserversstr, _ := opts.String("<realservers>")
			var realservers []fortigate.VIPRealservers
			for _, rs := range strings.Split(realserversstr, ",") {
				parts := strings.Split(rs, ":")
				rip := parts[0]
				rport, err := strconv.Atoi(parts[1])
				if err != nil {
					log.Fatalf("%s", err.Error())
				}
				realservers = append(realservers, fortigate.VIPRealservers{Ip: rip, Port: rport})
			}
			vip := &fortigate.VIP{
				Name:            name,
				Type:            fortigate.VIPTypeServerLoadBalance,
				LdbMethod:       fortigate.VIPLdbMethodRoundRobin,
				PortmappingType: fortigate.VIPPortmappingType1To1,
				Extintf:         "any",
				ServerType:      fortigate.VIPServerTypeTcp,
				Extip:           extip,
				Extport:         extport,
				Realservers:     realservers,
			}

			err = c.CreateVIP(vip)
			if err != nil {
				log.Fatalf("%s", err.Error())
			}
		}
	}
}
