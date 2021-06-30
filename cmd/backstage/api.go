package backstage

type API struct {
	ApiVersion string   `yaml:"apiVersion" json:"apiVersion"`
	Kind       string   `yaml:"kind" json:"kind"`
	Metadata   Metadata `yaml:"metadata" json:"metadata"`
	Spec       Spec     `yaml:"spec" json:"spec"`
}

// Metadata
type Metadata struct {
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
}

// Spec
type Spec struct {
	Type       string `yaml:"type" json:"type"`
	Lifecycle  string `yaml:"lifecycle" json:"lifecycle"`
	Owner      string `yaml:"owner" json:"owner"`
	System     string `yaml:"system" json:"system"`
	Definition string `yaml:"definition" json:"definition"`
}
