package hue

import "strings"

func HueWebhookHandler(data string) {

	if data == "all-off" {
		SetAllLightsState(false)
		return
	}

	if data == "strobe" {
		BlinkAll()
		return
	}

	if strings.HasPrefix(data, "to-hex-color:") {
		hex := strings.Replace(data, "to-hex-color:", "", 1)
		SetALlLightsColor(hex)
		return
	}

}
