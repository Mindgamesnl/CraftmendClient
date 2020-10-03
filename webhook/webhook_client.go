package webhook

import (
	"fmt"
	"github.com/Mindgamesnl/CraftmendClient/config"
	"github.com/Mindgamesnl/CraftmendWebhookClient"
)

func InitializeWebhook() {
	client := CraftmendWebhookClient.CreateWebhookClient(config.LoadedConfig.WebhookApiKey)

	client.On(func(event string, data string) {
		fmt.Println(event, " ", data)
	})

}
