package cmd

type Config struct {
	DeLijnAPIKey string  `yaml:"delijn_api_key"`
	Aliases      []Alias `yaml:"aliases"`
}

type Alias struct {
	Name     string          `yaml:"name"`
	Provider TransitProvider `yaml:"provider"`
	ID       []string        `yaml:"ID"`
}
