package main

import (
	"github.com/Mindgamesnl/CraftmendClient/config"
	"github.com/Mindgamesnl/CraftmendClient/hue"
	"github.com/Mindgamesnl/CraftmendClient/webhook"
	"github.com/sirupsen/logrus"
)

func main() {
	// load config
	config.LoadConfiguration()
	webhook.InitializeWebhook()
	hue.Init()

	logrus.Info("Booted.")

	// do fuck all
	select {

	}
}
