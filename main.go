package main

import (
	"github.com/Mindgamesnl/CraftmendClient/config"
	"github.com/Mindgamesnl/CraftmendClient/webhook"
)

func main() {
	// load config
	config.LoadConfiguration()
	webhook.InitializeWebhook()

	// do fuck all
	select {

	}
}
