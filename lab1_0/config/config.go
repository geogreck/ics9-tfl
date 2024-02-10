package config

type InputConfig struct {
	N        int    `yaml:"n"`
	Word     string `yaml:"word"`
	RawRules string `yaml:"rules"`
}
