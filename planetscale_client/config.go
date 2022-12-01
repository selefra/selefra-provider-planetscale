package planetscale_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

type Config struct {
	Token        string `yaml:"token,omitempty" mapstructure:"token"`
	Organization string `yaml:"organization,omitempty" mapstructure:"organization"`
}
