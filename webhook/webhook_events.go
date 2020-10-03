package webhook

import "github.com/Mindgamesnl/CraftmendClient/notifications"

func initializeHandlers(handler *eventHandler)  {
	handler.On("notify", notifications.FromWebhook)
}
