package config

type InputConfig struct {
	N         int      `yaml:"n"`
	Variables []string `yaml:"variables"`
	Word      string   `yaml:"word"`
	RawRules  string   `yaml:"rules"`
}
