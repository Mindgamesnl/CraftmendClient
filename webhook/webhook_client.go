package webhook

import (
	"github.com/Mindgamesnl/CraftmendClient/config"
	"github.com/Mindgamesnl/CraftmendWebhookClient"
	"github.com/sirupsen/logrus"
)

type eventHandler struct {
	handlers map[string][]func(data string)
}

func (eh *eventHandler) On(event string, handler func(data string)) {
	collection, found := eh.handlers[event]
	if !found {
		collection = []func(data string){}
	}
	collection = append(collection, handler)

	eh.handlers[event] = collection
}

func (eh eventHandler) do(event string, data string) {
	collection, found := eh.handlers[event]
	if found {
		for i := range collection {
			collection[i](data)
		}
	} else {
		logrus.Warn("No event handler for ", event)
	}
}

func InitializeWebhook() {
	client := CraftmendWebhookClient.CreateWebhookClient(config.LoadedConfig.WebhookApiKey)

	eh := eventHandler{
		handlers: make(map[string][]func(data string)),
	}

	initializeHandlers(&eh)

	client.On(func(event string, data string) {
		eh.do(event, data)
	})
}
