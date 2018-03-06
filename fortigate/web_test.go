package fortigate

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FortigateTestSuite struct {
	Vip string

	suite.Suite
}

func (s *FortigateTestSuite) TestVIP() {
	var vipresults VIPResults

	err := json.Unmarshal([]byte(s.Vip), &vipresults)
	assert.Nil(s.T(), err)

	if assert.Equal(s.T(), 1, len(vipresults.Results)) {
		vip := vipresults.Results[0]
		//spew.Dump(vip)
		assert.Equal(s.T(), "myvip", vip.Name)
		assert.Equal(s.T(), "10.70.125.1", vip.Extip)
		assert.Equal(s.T(), "443", vip.Extport)
		assert.Equal(s.T(), 2, len(vip.Monitor))
	}

}

func TestV546(t *testing.T) {
	j := `
{
	"http_method":"GET",
	"results":[
		{
			"name":"myvip",
			"q_origin_key":"myvip",
			"id":0,
			"uuid":"2452352352",
			"comment":"",
			"type":"server-load-balance",
			"dns-mapping-ttl":0,
			"ldb-method":"round-robin",
			"src-filter":[
			],
			"extip":"10.70.125.1",
			"mappedip":[
			],
			"mapped-addr":"",
			"extintf":"any",
			"arp-reply":"enable",
			"server-type":"tcp",
			"persistence":"none",
			"nat-source-vip":"disable",
			"portforward":"disable",
			"protocol":"tcp",
			"extport":"443",
			"mappedport":"0",
			"gratuitous-arp-interval":0,
			"srcintf-filter":[
			],
			"portmapping-type":"1-to-1",
			"realservers":[
				{
					"id":1,
					"q_origin_key":"1",
					"ip":"10.30.126.11",
					"port":443,
					"status":"active",
					"weight":1,
					"holddown-interval":300,
					"healthcheck":"vip",
					"http-host":"",
					"max-connections":0,
					"monitor":"",
					"client-ip":""
				},
				{
					"id":2,
					"q_origin_key":"2",
					"ip":"10.30.126.12",
					"port":443,
					"status":"active",
					"weight":1,
					"holddown-interval":300,
					"healthcheck":"vip",
					"http-host":"",
					"max-connections":0,
					"monitor":"",
					"client-ip":""
				},
				{
					"id":3,
					"q_origin_key":"3",
					"ip":"10.30.126.17",
					"port":443,
					"status":"active",
					"weight":1,
					"holddown-interval":300,
					"healthcheck":"vip",
					"http-host":"",
					"max-connections":0,
					"monitor":"",
					"client-ip":""
				}
			],
			"http-cookie-domain-from-host":"disable",
			"http-cookie-domain":"",
			"http-cookie-path":"",
			"http-cookie-generation":0,
			"http-cookie-age":60,
			"http-cookie-share":"same-ip",
			"https-cookie-secure":"disable",
			"http-multiplex":"disable",
			"http-ip-header":"disable",
			"http-ip-header-name":"",
			"outlook-web-access":"disable",
			"weblogic-server":"disable",
			"websphere-server":"disable",
			"ssl-mode":"half",
			"ssl-certificate":"",
			"ssl-dh-bits":"2048",
			"ssl-algorithm":"high",
			"ssl-cipher-suites":[
			],
			"ssl-server-algorithm":"client",
			"ssl-server-cipher-suites":[
			],
			"ssl-pfs":"allow",
			"ssl-min-version":"tls-1.0",
			"ssl-max-version":"tls-1.2",
			"ssl-server-min-version":"client",
			"ssl-server-max-version":"client",
			"ssl-send-empty-frags":"enable",
			"ssl-client-fallback":"enable",
			"ssl-client-renegotiation":"allow",
			"ssl-client-session-state-type":"both",
			"ssl-client-session-state-timeout":30,
			"ssl-client-session-state-max":1000,
			"ssl-server-session-state-type":"both",
			"ssl-server-session-state-timeout":60,
			"ssl-server-session-state-max":100,
			"ssl-http-location-conversion":"disable",
			"ssl-http-match-host":"disable",
			"monitor":"\"CheckTCP\" \"CheckTCP2\"",
			"max-embryonic-connections":1000,
			"color":0
		}
	],
	"vdom":"MYVDOM",
	"path":"firewall",
	"name":"vip",
	"mkey":"myvip",
	"status":"success",
	"http_status":200,
	"serial":"FG10223DSF34df",
	"version":"v5.4.6",
	"build":6408
}`
	s := &FortigateTestSuite{
		Vip: j,
	}

	apiVersion = "v5.4.6"

	suite.Run(t, s)
}

func TestV563(t *testing.T) {
	j := `
{
  "http_method":"GET",
  "revision":"1958.0.82.573477330.1513174867",
  "results":[
    {
      "name":"myvip",
      "q_origin_key":"myvip",
      "id":0,
      "uuid":"24523523625234",
      "comment":"mycomment",
      "type":"server-load-balance",
      "dns-mapping-ttl":0,
      "ldb-method":"round-robin",
      "src-filter":[
      ],
      "service":[
      ],
      "extip":"10.70.125.1",
      "extaddr":[
      ],
      "mappedip":[
      ],
      "mapped-addr":"",
      "extintf":"any",
      "arp-reply":"enable",
      "server-type":"tcp",
      "persistence":"none",
      "nat-source-vip":"disable",
      "portforward":"disable",
      "protocol":"tcp",
      "extport":"443",
      "mappedport":"0",
      "gratuitous-arp-interval":0,
      "srcintf-filter":[
      ],
      "portmapping-type":"1-to-1",
      "realservers":[
        {
          "id":1,
          "q_origin_key":1,
          "ip":"10.30.126.11",
          "port":443,
          "status":"active",
          "weight":1,
          "holddown-interval":300,
          "healthcheck":"vip",
          "http-host":"",
          "max-connections":0,
          "monitor":"",
          "client-ip":""
        },
        {
          "id":2,
          "q_origin_key":2,
          "ip":"10.30.126.12",
          "port":443,
          "status":"active",
          "weight":1,
          "holddown-interval":300,
          "healthcheck":"vip",
          "http-host":"",
          "max-connections":0,
          "monitor":"",
          "client-ip":""
        },
        {
          "id":3,
          "q_origin_key":3,
          "ip":"10.30.126.13",
          "port":443,
          "status":"active",
          "weight":1,
          "holddown-interval":300,
          "healthcheck":"vip",
          "http-host":"",
          "max-connections":0,
          "monitor":"",
          "client-ip":""
        }
      ],
      "http-cookie-domain-from-host":"disable",
      "http-cookie-domain":"",
      "http-cookie-path":"",
      "http-cookie-generation":0,
      "http-cookie-age":60,
      "http-cookie-share":"same-ip",
      "http-multiplex":"disable",
      "http-ip-header":"disable",
      "http-ip-header-name":"",
      "outlook-web-access":"disable",
      "weblogic-server":"disable",
      "websphere-server":"disable",
      "monitor":[
        {
          "name":"CheckTCP",
          "q_origin_key":"CheckTCP"
        },
        {
          "name":"CheckTCP2",
          "q_origin_key":"CheckTCP2"
        }
      ],
      "max-embryonic-connections":1000,
      "color":0
    }
  ],
  "vdom":"root",
  "path":"firewall",
  "name":"vip",
  "mkey":"myvip",
  "status":"success",
  "http_status":200,
  "serial":"FG10223DSF34df",
  "version":"v5.6.3",
  "build":1547
}
`
	s := &FortigateTestSuite{
		Vip: j,
	}

	apiVersion = "v5.6.3"

	suite.Run(t, s)
}
