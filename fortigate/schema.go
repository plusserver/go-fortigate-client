package fortigate

import "net/url"

type SchemaResponse struct {
	Endpoints []Endpoint `json:"results"`
	Result
}

type Endpoint struct {
	Path   string `json:"path"`
	Name   string `json:"name"`
	Alias  string
	Schema Schema `json:"schema"`
}

type Schema struct {
	Name     string                 `json:"name"`
	Category string                 `json:"category"`
	Mkey     string                 `json:"mkey"`
	MkeyType string                 `json:"mkey_type",omitempty`
	Help     string                 `json:"help"`
	Children map[string]SchemaChild `json:"children"`
}

type SchemaChild struct {
	Name     string                 `json:"name",omitempty`
	Category string                 `json:"category",omitempty`
	Type     string                 `json:"type",omitempty`
	Help     string                 `json:"help",omitempty`
	Options  []SchemaOption         `json:"options",omitempty`
	Children map[string]SchemaChild `json:"children",omitempty`
}

type SchemaOption struct {
	Name string `json:"name",omitempty"`
	Help string `json:"help",omitempty"`
}

func (c *WebClient) Schema() ([]Endpoint, error) {
	var resp SchemaResponse

	_, err := c.napping.Get(c.URL+"/api/v2/cmdb", &url.Values{"action": []string{"schema"}}, &resp, nil)
	if err != nil {
		return []Endpoint{}, err
	}
	return resp.Endpoints, nil
}
