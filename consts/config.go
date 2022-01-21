package consts

import "os"

const (
	GanacheCLI string = "ganache-cli"
	BSCTest    string = "bsctest"
)

type IConfig interface {
	Network() string
}

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) Network() string {
	return os.Getenv("NETWORK")
}
