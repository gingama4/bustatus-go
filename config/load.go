package config

import (
	"encoding/json"
	"io/ioutil"
)

func initConfig() *Config {
	c := Config{
		Format: "{{.statusEmoji}} {{.artist}} - {{.title}}",
		Emoji: &Emoji{
			Play:       "",
			Pause:      "",
			Stop:       "",
			Connect:    "",
			NotConnect: "",
		},
		NotConnMes: "{{.statusEmoji}}",
	}

	return &c
}

func GetConfig(filepath string) (*Config, error) {
	c := initConfig()

	u, err := getConfig(filepath)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return c, nil
	}

	if u.Format != "" {
		c.Format = u.Format
	}

	if u.NotConnMes != "" {
		c.NotConnMes = u.NotConnMes
	}

	if u.Emoji != nil {
		c.Emoji = u.Emoji
	}

	return c, nil
}

func getConfig(filepath string) (*Config, error) {
	u := Config{}
	if filepath == "" {
		return nil, nil
	}

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &u); err != nil {
		return nil, err
	}

	return &u, nil
}
