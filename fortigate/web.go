package fortigate

import (
	"fmt"
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"gopkg.in/jmcvetta/napping.v3"
)

const (
	CsrfToken = "ccsrftoken"
)

type WebClient struct {
	URL      string
	User     string
	Password string
	ApiKey   string
	Log      bool

	napping *napping.Session
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

func NewWebClient(clientConfig WebClient) (c *WebClient, err error) {

	if clientConfig.URL == "" {
		return nil, fmt.Errorf("URL is required")
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}

	session := &napping.Session{
		Client: client,
		Log:    clientConfig.Log,
	}

	c = &WebClient{
		URL:      strings.TrimRight(clientConfig.URL, "/"),
		User:     clientConfig.User,
		Password: clientConfig.Password,
		ApiKey:   clientConfig.ApiKey,
		napping:  session,
	}

	if clientConfig.ApiKey != "" {
		c.napping.Header = &http.Header{
			"Authorization": []string{"Bearer " + clientConfig.ApiKey},
		}
	} else if clientConfig.User != "" && clientConfig.Password != "" {
		err = c.authenticate()
		if err != nil {
			return
		}
	} else {
		return nil, fmt.Errorf("either an API key or a user/password is required")
	}

	return
}

func (c *WebClient) authenticate() error {
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return fmt.Errorf("could not create cookie jar for fortigate web client: %s", err.Error())
	}
	c.napping.Client.Jar = jar

	resp, err := c.napping.Client.PostForm(c.URL+"/logincheck", url.Values{"username": {c.User}, "secretkey": {c.Password}})
	if err != nil {
		return fmt.Errorf("could not log in user '%s' at '%s': %s", c.User, c.URL, err.Error())
	}

	for _, cookie := range resp.Cookies() {
		if cookie.Name == CsrfToken {
			if cookie.Value == "0%260" {
				return fmt.Errorf("could not log in as user '%s' at '%s'", c.User, c.URL)
			} else {
				c.napping.Header = &http.Header{"X-CSRFTOKEN": []string{cookie.Value}}
			}
		}
	}

	return nil
}

func (c *WebClient) do(method string, path string, p *url.Values, payload interface{}, result2 interface{}) (resp *napping.Response, err error) {
	var result interface{}
	var errMsg interface{}

	path = c.URL + "/api/v2/cmdb/" + path

	switch method {
	case http.MethodGet:
		result = result2
		resp, err = c.napping.Get(path, p, result, errMsg)
	case http.MethodDelete:
		resp, err = c.napping.Delete(path, p, result, errMsg)
	case http.MethodPost:
		resp, err = c.napping.Post(path, payload, result, errMsg)
	case http.MethodPut:
		resp, err = c.napping.Put(path, payload, result, errMsg)
	}

	if err != nil {
		return
	}

	if resp.HttpResponse().StatusCode == 401 {
		err = c.authenticate()
		if err != nil {
			return
		}
		switch method {
		case http.MethodGet:
			result = result2
			resp, err = c.napping.Get(path, p, result, errMsg)
		case http.MethodDelete:
			resp, err = c.napping.Delete(path, p, result, errMsg)
		case http.MethodPost:
			resp, err = c.napping.Post(path, payload, result, errMsg)
		case http.MethodPut:
			resp, err = c.napping.Put(path, payload, result, errMsg)
		}
		if err != nil {
			return
		}
		if resp.HttpResponse().StatusCode == 401 {
			return nil, fmt.Errorf("authentication failed")
		}
	}

	if rres, ok := result.(*Result); ok {
		if rres.HTTPStatus != 200 {
			if rerrmsg, ok := errMsg.(*Result); ok {
				if rerrmsg.HTTPStatus == 404 {
					err = fmt.Errorf("not found")
				} else if rerrmsg.HTTPStatus == 401 {
					err = fmt.Errorf("unauthorized")
				} else {
					err = fmt.Errorf(rerrmsg.Status)
				}
			}
		}
	}

	return
}
