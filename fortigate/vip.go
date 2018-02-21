// WARNING: This file was generated by gen/generator.go

package fortigate

import (
	"fmt"
)

// Enable to respond to ARP requests for this virtual IP address. Enabled by default.
type VIPArpReply string

// External FQDN address name.
type VIPExtaddr struct {

	// Address name.
	Name string `json:"name,omitempty"`
}

// Enable/disable use of HTTP cookie domain from host field in HTTP.
type VIPHttpCookieDomainFromHost string

// Control sharing of cookies across virtual servers. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing.
type VIPHttpCookieShare string

// For HTTP multiplexing, enable to add the original client IP address in the XForwarded-For HTTP header.
type VIPHttpIpHeader string

// Enable/disable HTTP multiplexing.
type VIPHttpMultiplex string

// Method used to distribute sessions to real servers.
type VIPLdbMethod string

// IP address or address range on the destination network to which the external IP address is mapped.
type VIPMappedip struct {

	// Mapped IP range.
	Range string `json:"range,omitempty"`
}

// Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
type VIPMonitor struct {

	// Health monitor name.
	Name string `json:"name,omitempty"`
}

// Enable to prevent unintended servers from using a virtual IP. Disable to use the actual IP address of the server as the source address.
type VIPNatSourceVip string

// Enable to add the Front-End-Https header for Microsoft Outlook Web Access.
type VIPOutlookWebAccess string

// Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session.
type VIPPersistence string

// Enable/disable port forwarding.
type VIPPortforward string

// Port mapping type.
type VIPPortmappingType string

// Protocol to use when forwarding packets.
type VIPProtocol string

// Select the real servers that this server load balancing VIP will distribute traffic to.
type VIPRealservers struct {

	// Only clients in this IP range can connect to this real server.
	ClientIp string `json:"client-ip,omitempty"`

	// Enable to check the responsiveness of the real server before forwarding traffic.
	Healthcheck string `json:"healthcheck,omitempty"`

	// Time in seconds that the health check monitor continues to monitor and unresponsive server that should be active.
	HolddownInterval int `json:"holddown-interval,omitempty"`

	// HTTP server domain name in HTTP header.
	HttpHost string `json:"http-host,omitempty"`

	// Real server ID.
	Id int `json:"id,omitempty"`

	// IP address of the real server.
	Ip string `json:"ip,omitempty"`

	// Max number of active connections that can be directed to the real server. When reached, sessions are sent to other real servers.
	MaxConnections int `json:"max-connections,omitempty"`

	// Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
	Monitor string `json:"monitor,omitempty"`

	// Port for communicating with the real server. Required if port forwarding is enabled.
	Port int `json:"port,omitempty"`

	// Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent.
	Status string `json:"status,omitempty"`

	// Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
	Weight int `json:"weight,omitempty"`
}

// Protocol to be load balanced by the virtual server (also called the server load balance virtual IP).
type VIPServerType string

// Service name.
type VIPService struct {

	// Service name.
	Name string `json:"name,omitempty"`
}

// Source address filter. Each address must be either an IP/subnet (x.x.x.x/n) or a range (x.x.x.x-y.y.y.y). Separate addresses with spaces.
type VIPSrcFilter struct {

	// Source-filter range.
	Range string `json:"range,omitempty"`
}

// Interfaces to which the VIP applies. Separate the names with spaces.
type VIPSrcintfFilter struct {

	// Interface name.
	InterfaceName string `json:"interface-name,omitempty"`
}

// Configure a static NAT, load balance, server load balance, DNS translation, or FQDN VIP.
type VIPType string

// Enable to add an HTTP header to indicate SSL offloading for a WebLogic server.
type VIPWeblogicServer string

// Enable to add an HTTP header to indicate SSL offloading for a WebSphere server.
type VIPWebsphereServer string

const (

	// Disable ARP reply.
	VIPArpReplyDisable VIPArpReply = "disable"

	// Enable ARP reply.
	VIPArpReplyEnable VIPArpReply = "enable"

	// Disable use of HTTP cookie domain from host field in HTTP (use http-cooke-domain setting).
	VIPHttpCookieDomainFromHostDisable VIPHttpCookieDomainFromHost = "disable"

	// Enable use of HTTP cookie domain from host field in HTTP.
	VIPHttpCookieDomainFromHostEnable VIPHttpCookieDomainFromHost = "enable"

	// Only allow HTTP cookie to match this virtual server.
	VIPHttpCookieShareDisable VIPHttpCookieShare = "disable"

	// Allow HTTP cookie to match any virtual server with same IP.
	VIPHttpCookieShareSameIp VIPHttpCookieShare = "same-ip"

	// Enable adding HTTP header.
	VIPHttpIpHeaderEnable VIPHttpIpHeader = "enable"

	// Disable adding HTTP header.
	VIPHttpIpHeaderDisable VIPHttpIpHeader = "disable"

	// Enable HTTP session multiplexing.
	VIPHttpMultiplexEnable VIPHttpMultiplex = "enable"

	// Disable HTTP session multiplexing.
	VIPHttpMultiplexDisable VIPHttpMultiplex = "disable"

	// Distribute to server based on source IP.
	VIPLdbMethodStatic VIPLdbMethod = "static"

	// Distribute to server based round robin order.
	VIPLdbMethodRoundRobin VIPLdbMethod = "round-robin"

	// Distribute to server based on weight.
	VIPLdbMethodWeighted VIPLdbMethod = "weighted"

	// Distribute to server with lowest session count.
	VIPLdbMethodLeastSession VIPLdbMethod = "least-session"

	// Distribute to server with lowest Round-Trip-Time.
	VIPLdbMethodLeastRtt VIPLdbMethod = "least-rtt"

	// Distribute to the first server that is alive.
	VIPLdbMethodFirstAlive VIPLdbMethod = "first-alive"

	// Distribute to server based on host field in HTTP header.
	VIPLdbMethodHttpHost VIPLdbMethod = "http-host"

	// Do not force to NAT as VIP.
	VIPNatSourceVipDisable VIPNatSourceVip = "disable"

	// Force to NAT as VIP.
	VIPNatSourceVipEnable VIPNatSourceVip = "enable"

	// Disable Outlook Web Access support.
	VIPOutlookWebAccessDisable VIPOutlookWebAccess = "disable"

	// Enable Outlook Web Access support.
	VIPOutlookWebAccessEnable VIPOutlookWebAccess = "enable"

	// None.
	VIPPersistenceNone VIPPersistence = "none"

	// HTTP cookie.
	VIPPersistenceHttpCookie VIPPersistence = "http-cookie"

	// Disable port forward.
	VIPPortforwardDisable VIPPortforward = "disable"

	// Enable port forward.
	VIPPortforwardEnable VIPPortforward = "enable"

	// One to one.
	VIPPortmappingType1To1 VIPPortmappingType = "1-to-1"

	// Many to many.
	VIPPortmappingTypeMToN VIPPortmappingType = "m-to-n"

	// TCP.
	VIPProtocolTcp VIPProtocol = "tcp"

	// UDP.
	VIPProtocolUdp VIPProtocol = "udp"

	// SCTP.
	VIPProtocolSctp VIPProtocol = "sctp"

	// ICMP.
	VIPProtocolIcmp VIPProtocol = "icmp"

	// HTTP
	VIPServerTypeHttp VIPServerType = "http"

	// TCP
	VIPServerTypeTcp VIPServerType = "tcp"

	// UDP
	VIPServerTypeUdp VIPServerType = "udp"

	// IP
	VIPServerTypeIp VIPServerType = "ip"

	// Static NAT.
	VIPTypeStaticNat VIPType = "static-nat"

	// Load balance.
	VIPTypeLoadBalance VIPType = "load-balance"

	// Server load balance.
	VIPTypeServerLoadBalance VIPType = "server-load-balance"

	// DNS translation.
	VIPTypeDnsTranslation VIPType = "dns-translation"

	// Fully qualified domain name.
	VIPTypeFqdn VIPType = "fqdn"

	// Do not add HTTP header indicating SSL offload for WebLogic server.
	VIPWeblogicServerDisable VIPWeblogicServer = "disable"

	// Add HTTP header indicating SSL offload for WebLogic server.
	VIPWeblogicServerEnable VIPWeblogicServer = "enable"

	// Do not add HTTP header indicating SSL offload for WebSphere server.
	VIPWebsphereServerDisable VIPWebsphereServer = "disable"

	// Add HTTP header indicating SSL offload for WebSphere server.
	VIPWebsphereServerEnable VIPWebsphereServer = "enable"
)

// Configure virtual IP for IPv4.
type VIP struct {

	// Enable to respond to ARP requests for this virtual IP address. Enabled by default.
	ArpReply VIPArpReply `json:"arp-reply,omitempty"`

	// Color of icon on the GUI.
	Color int `json:"color,omitempty"`

	// Comment.
	Comment string `json:"comment,omitempty"`

	// DNS mapping TTL (Set to zero to use TTL in DNS response, default = 0).
	DnsMappingTtl int `json:"dns-mapping-ttl,omitempty"`

	// External FQDN address name.
	Extaddr []VIPExtaddr `json:"extaddr,omitempty"`

	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	Extintf string `json:"extintf,omitempty"`

	// IP address or address range on the external interface that you want to map to an address or address range on the destination network.
	Extip string `json:"extip,omitempty"`

	// Incoming port number range that you want to map to a port number range on the destination network.
	Extport string `json:"extport,omitempty"`

	// Enable to have the VIP send gratuitous ARPs. 0=disabled. Set from 5 up to 8640000 seconds to enable.
	GratuitousArpInterval int `json:"gratuitous-arp-interval,omitempty"`

	// Time in minutes that client web browsers should keep a cookie. Default is 60 seconds. 0 = no time limit.
	HttpCookieAge int `json:"http-cookie-age,omitempty"`

	// Domain that HTTP cookie persistence should apply to.
	HttpCookieDomain string `json:"http-cookie-domain,omitempty"`

	// Enable/disable use of HTTP cookie domain from host field in HTTP.
	HttpCookieDomainFromHost VIPHttpCookieDomainFromHost `json:"http-cookie-domain-from-host,omitempty"`

	// Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
	HttpCookieGeneration int `json:"http-cookie-generation,omitempty"`

	// Limit HTTP cookie persistence to the specified path.
	HttpCookiePath string `json:"http-cookie-path,omitempty"`

	// Control sharing of cookies across virtual servers. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing.
	HttpCookieShare VIPHttpCookieShare `json:"http-cookie-share,omitempty"`

	// For HTTP multiplexing, enable to add the original client IP address in the XForwarded-For HTTP header.
	HttpIpHeader VIPHttpIpHeader `json:"http-ip-header,omitempty"`

	// For HTTP multiplexing, enter a custom HTTPS header name. The original client IP address is added to this header. If empty, X-Forwarded-For is used.
	HttpIpHeaderName string `json:"http-ip-header-name,omitempty"`

	// Enable/disable HTTP multiplexing.
	HttpMultiplex VIPHttpMultiplex `json:"http-multiplex,omitempty"`

	// Custom defined ID.
	Id int `json:"id,omitempty"`

	// Method used to distribute sessions to real servers.
	LdbMethod VIPLdbMethod `json:"ldb-method,omitempty"`

	// Mapped FQDN address name.
	MappedAddr string `json:"mapped-addr,omitempty"`

	// IP address or address range on the destination network to which the external IP address is mapped.
	Mappedip []VIPMappedip `json:"mappedip,omitempty"`

	// Port number range on the destination network to which the external port number range is mapped.
	Mappedport string `json:"mappedport,omitempty"`

	// Maximum number of incomplete connections.
	MaxEmbryonicConnections int `json:"max-embryonic-connections,omitempty"`

	// Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
	Monitor []VIPMonitor `json:"monitor,omitempty"`

	// Virtual IP name.
	Name string `json:"name,omitempty"`

	// Enable to prevent unintended servers from using a virtual IP. Disable to use the actual IP address of the server as the source address.
	NatSourceVip VIPNatSourceVip `json:"nat-source-vip,omitempty"`

	// Enable to add the Front-End-Https header for Microsoft Outlook Web Access.
	OutlookWebAccess VIPOutlookWebAccess `json:"outlook-web-access,omitempty"`

	// Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session.
	Persistence VIPPersistence `json:"persistence,omitempty"`

	// Enable/disable port forwarding.
	Portforward VIPPortforward `json:"portforward,omitempty"`

	// Port mapping type.
	PortmappingType VIPPortmappingType `json:"portmapping-type,omitempty"`

	// Protocol to use when forwarding packets.
	Protocol VIPProtocol `json:"protocol,omitempty"`

	// Select the real servers that this server load balancing VIP will distribute traffic to.
	Realservers []VIPRealservers `json:"realservers,omitempty"`

	// Protocol to be load balanced by the virtual server (also called the server load balance virtual IP).
	ServerType VIPServerType `json:"server-type,omitempty"`

	// Service name.
	Service []VIPService `json:"service,omitempty"`

	// Source address filter. Each address must be either an IP/subnet (x.x.x.x/n) or a range (x.x.x.x-y.y.y.y). Separate addresses with spaces.
	SrcFilter []VIPSrcFilter `json:"src-filter,omitempty"`

	// Interfaces to which the VIP applies. Separate the names with spaces.
	SrcintfFilter []VIPSrcintfFilter `json:"srcintf-filter,omitempty"`

	// Configure a static NAT, load balance, server load balance, DNS translation, or FQDN VIP.
	Type VIPType `json:"type,omitempty"`

	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid string `json:"uuid,omitempty"`

	// Enable to add an HTTP header to indicate SSL offloading for a WebLogic server.
	WeblogicServer VIPWeblogicServer `json:"weblogic-server,omitempty"`

	// Enable to add an HTTP header to indicate SSL offloading for a WebSphere server.
	WebsphereServer VIPWebsphereServer `json:"websphere-server,omitempty"`
}

// The results of a Get or List operation
type VIPResults struct {
	Results []*VIP `json:"results"`
	Mkey    string `json:"mkey"`
	Result
}

// List all VIPs
func (c *WebClient) ListVIPs() (res []*VIP, err error) {
	var errmsg Result
	var results VIPResults
	_, err = c.napping.Get(c.URL+"/api/v2/cmdb/firewall/vip", nil, &results, nil)
	if err != nil {
		return []*VIP{}, fmt.Errorf("error listing VIPs: %s", err.Error())
	}
	if results.HTTPStatus != 200 {
		if errmsg.HTTPStatus == 404 {
			return []*VIP{}, fmt.Errorf("error listing VIP: not found")
		} else {
			return []*VIP{}, fmt.Errorf("error listing VIP: %s", errmsg.Status)
		}
	}
	res = results.Results
	return
}

// Get a VIP by name
func (c *WebClient) GetVIP(name string) (res *VIP, err error) {
	var errmsg Result
	var results VIPResults
	_, err = c.napping.Get(c.URL+"/api/v2/cmdb/firewall/vip/"+name, nil, &results, &errmsg)
	if err != nil {
		return &VIP{}, fmt.Errorf("error getting VIP '%s': %s", name, err.Error())
	}
	if results.HTTPStatus != 200 {
		if errmsg.HTTPStatus == 404 {
			return &VIP{}, fmt.Errorf("error getting VIP '%s': not found", name)
		} else {
			return &VIP{}, fmt.Errorf("error getting VIP '%s': %s", name, errmsg.Status)
		}
	}
	if len(results.Results) == 0 {
		return &VIP{}, fmt.Errorf("error getting VIP '%s': not found", name)
	}

	res = results.Results[0]
	return
}

// Create a new VIP
func (c *WebClient) CreateVIP(obj *VIP) (id string, err error) {
	var errmsg Result
	var results VIPResults
	_, err = c.napping.Post(c.URL+"/api/v2/cmdb/firewall/vip", obj, &results, &errmsg)
	if err != nil {
		return "", fmt.Errorf("error creating VIP '%s': %s", obj.Name, err.Error())
	}
	if results.HTTPStatus == 200 {
		return
	}
	if errmsg.HTTPStatus != 200 {
		return "", fmt.Errorf("error creating VIP '%s': %s", obj.Name, errmsg.Status)
	}

	return
}

// Update a VIP
func (c *WebClient) UpdateVIP(obj *VIP) (err error) {
	var errmsg Result
	var results VIPResults
	_, err = c.napping.Put(c.URL+"/api/v2/cmdb/firewall/vip/"+obj.Name, obj, &results, &errmsg)
	if err != nil {
		return fmt.Errorf("error updating VIP '%s': %s", obj.Name, err.Error())
	}
	if results.HTTPStatus != 200 {
		if errmsg.HTTPStatus == 404 {
			return fmt.Errorf("error updating VIP '%s': not found", obj.Name)
		} else {
			return fmt.Errorf("error updating VIP '%s': %s", obj.Name, errmsg.Status)
		}
	}

	return
}

// Delete a VIP by name
func (c *WebClient) DeleteVIP(name string) (err error) {
	var errmsg Result
	var results VIPResults
	_, err = c.napping.Delete(c.URL+"/api/v2/cmdb/firewall/vip/"+name, nil, &results, &errmsg)
	if err != nil {
		return fmt.Errorf("error deleting VIP '%s': %s", name, err.Error())
	}
	if results.HTTPStatus == 200 {
		return
	}
	if errmsg.HTTPStatus != 200 {
		if errmsg.HTTPStatus == 404 {
			return fmt.Errorf("error deleting VIP '%s': not found", name)
		}
		return fmt.Errorf("error deleting VIP '%s': %s", name, errmsg.Status)
	}

	return
}

// List all VIPs
func (c *FakeClient) ListVIPs() (res []*VIP, err error) {
	for _, r := range c.VIPs {
		res = append(res, r)
	}
	return
}

// Get a VIP by name
func (c *FakeClient) GetVIP(name string) (*VIP, error) {
	if res, ok := c.VIPs[name]; ok {
		return res, nil
	} else {
		return &VIP{}, fmt.Errorf("error getting VIP '%s': not found", name)
	}
}

// Create a new VIP
func (c *FakeClient) CreateVIP(obj *VIP) (id string, err error) {
	c.VIPs[obj.Name] = obj
	return "", nil
}

// Update a VIP
func (c *FakeClient) UpdateVIP(obj *VIP) (err error) {
	c.VIPs[obj.Name] = obj
	return nil
}

// Delete a VIP by name
func (c *FakeClient) DeleteVIP(name string) (err error) {
	delete(c.VIPs, name)
	return nil
}
