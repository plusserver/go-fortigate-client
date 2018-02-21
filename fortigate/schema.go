package fortigate

import "net/url"

type SchemaResults struct {
	Name     string                 `json:"name"`
	Category string                 `json:"category"`
	MkeyType string                 `json:"mkey_type",omitempty`
	Help     string                 `json:"help"`
	Children map[string]SchemaChild `json:"children"`
}

type Schema struct {
	TypeName string
	Path     string
	Results  SchemaResults `json:"results",omitempty`
	Result
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

func (c *WebClient) Schema(path string) (schema Schema) {
	_, err := c.napping.Get(c.URL+"/api/v2"+path, &url.Values{"action": []string{"schema"}}, &schema, nil)
	if err != nil {
		panic(err)
	}
	return
}
