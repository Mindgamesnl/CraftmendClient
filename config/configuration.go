package config

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

type Config struct {
	WebhookApiKey string `json:"webhook-api-key"`
	HueUser string `json:"hue-user"`
}

var LoadedConfig Config

var defaultConfig = `
{
	"webhook-api-key": "craftmend webhook api key",
	"hue-user": "hue user key"
}
`

func LoadConfiguration() Config {
	f, err := os.Open("./config.json")
	if err != nil {
		logrus.Info("Creating default config")
		_ = ioutil.WriteFile("./config.json", []byte(defaultConfig), 0644)
		return LoadConfiguration()
	}
	defer f.Close()

	logrus.Info("Loading config")

	var cfg Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		logrus.Error(err)
	}

	LoadedConfig = cfg
	return cfg
}
