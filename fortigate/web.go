package fortigate

import (
	"crypto/tls"
	"net/http"
	"strings"

	"gopkg.in/jmcvetta/napping.v3"
)

type WebClient struct {
	URL      string
	User     string
	Password string
	ApiKey   string
	Log      bool

	napping napping.Session
}

type Result struct {
	HTTPMethod string `json:"http_method"`
	Revision   string `json:"revision"`
	Status     string `json:"status"`
	HTTPStatus int    `json:"http_status"`
	Vdom       string `json:"vdom"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	Serial     string `json:"serial"`
	Version    string `json:"version"`
	Build      int    `json:"build"`
	Action     string `json:"action"`
}

func NewWebClient(c WebClient) *WebClient {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}

	session := napping.Session{
		Client: client,
		Log:    c.Log,
	}

	if c.ApiKey != "" {
		session.Header = &http.Header{
			"Authorization": []string{"Bearer " + c.ApiKey},
		}
	}

	return &WebClient{
		URL:      strings.TrimRight(c.URL, "/"),
		User:     c.User,
		Password: c.Password,
		ApiKey:   c.ApiKey,
		napping:  session,
	}
}
