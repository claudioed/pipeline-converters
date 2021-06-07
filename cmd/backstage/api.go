package backstage

type API struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

// Metadata
type Metadata struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

// Spec
type Spec struct {
	Type       string `yaml:"type"`
	Lifecycle  string `yaml:"lifecycle"`
	Owner      string `yaml:"owner"`
	System     string `yaml:"system"`
	Definition string `yaml:"definition"`
}
