package fortigate

import (
	"fmt"
	"net/url"
)

type SchemaResponse struct {
	Endpoints []Endpoint `json:"results,omitempty"`
	Result
}

type Endpoint struct {
	Path   string `json:"path"`
	Name   string `json:"name"`
	Alias  string
	Schema Schema `json:"schema"`
}

type Schema struct {
	Name     string                 `json:"name,omitempty"`
	Category string                 `json:"category,omitempty"`
	Mkey     string                 `json:"mkey,omitempty"`
	MkeyType string                 `json:"mkey_type,omitempty"`
	Help     string                 `json:"help,omitempty"`
	Children map[string]SchemaChild `json:"children,omitempty"`
}

type SchemaChild struct {
	Name     string                 `json:"name,omitempty"`
	Category string                 `json:"category,omitempty"`
	Type     string                 `json:"type,omitempty"`
	Help     string                 `json:"help,omitempty"`
	Options  []SchemaOption         `json:"options,omitempty"`
	Children map[string]SchemaChild `json:"children,omitempty"`
}

type SchemaOption struct {
	Name string `json:"name,omitempty"`
	Help string `json:"help,omitempty"`
}

func (c *WebClient) Schema() ([]Endpoint, error) {
	var resp SchemaResponse

	_, err := c.napping.Get(c.URL+"/api/v2/cmdb", &url.Values{"action": []string{"schema"}}, &resp, nil)

	if err != nil {
		fmt.Println("error!")
		return []Endpoint{}, err
	}
	return resp.Endpoints, nil
}
