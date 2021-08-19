package backstage

type Location struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Spec struct {
		Type    string   `json:"type"`
		Targets []string `json:"targets"`
	} `json:"spec"`
}
