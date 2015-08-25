package main

import (
	"github.com/BurntSushi/toml"
	"github.com/operando/golack"
)

type Config struct {
	Slack     golack.Slack   `toml:"slack"`
	Webhook   golack.Webhook `toml:"webhook"`
	Payload   golack.Payload `toml:"playload"`
	Log       string         `toml:"log"`
	Package   string         `toml:"package"`
	SleepTime int            `toml:"sleeptime"`
}

func LoadConfig(configPath string, config *Config) (*Config, error) {
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		return config, err
	}
	return config, nil
}
