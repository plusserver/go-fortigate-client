package fortigate

import (
	"crypto/tls"
	"net/http"

	"gopkg.in/jmcvetta/napping.v3"
	"strings"
)

type WebClient struct {
	URL      string
	User     string
	Password string
	ApiKey   string

	napping napping.Session
}

type Result struct {
	HTTPMethod string `json:"http_method"`
	Revision   string `json:"revision"`
	Mkey       string `json:"mkey"`
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
		//Log:    true,
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
