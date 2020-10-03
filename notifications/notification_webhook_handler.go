package notifications

import (
	"encoding/json"
	"github.com/martinlindhe/notify"
)

type alert struct {
	Title string `json:"title"`
	Body string `json:"body"`
	App string `json:"app"`
}

func FromWebhook(data string)  {
	var a alert
	json.Unmarshal([]byte(data), &a)
	notify.Alert(a.App, a.Title, a.Body, NOTIFY_OA_LOGO)
}
