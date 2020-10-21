package webhook

import (
	"github.com/Mindgamesnl/CraftmendClient/config"
	"github.com/Mindgamesnl/CraftmendWebhookClient"
	"github.com/Mindgamesnl/gjobs/jobs"
	"github.com/sirupsen/logrus"
)

var manager = jobs.NewJobManager(1)

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
		logrus.Info("Running ", event)
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
		logrus.Info("Scheduling " + event)
		manager.ScheduleFunction(func() {
			eh.do(event, data)
		})
	})
}
