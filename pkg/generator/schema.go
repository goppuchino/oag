package generator

type Spec struct {
	OpenAPI           string                   `json:"openapi" yaml:"openapi"`
	Info              map[string]interface{}   `json:"info" yaml:"info"`
	JsonSchemaDialect string                   `json:"jsonSchemaDialect,omitempty" yaml:"jsonSchemaDialect,omitempty"`
	Servers           []map[string]interface{} `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths             map[string]interface{}   `json:"paths" yaml:"paths"`
}

type PathSpec struct {
	Summary     string                   `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string                   `json:"description,omitempty" yaml:"description,omitempty"`
	Tags        []string                 `json:"tags,omitempty" yaml:"tags,omitempty"`
	Parameters  []map[string]interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
