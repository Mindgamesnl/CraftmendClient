package hue

import (
	"github.com/Mindgamesnl/CraftmendClient/config"
	"github.com/amimof/huego"
	"github.com/lucasb-eyer/go-colorful"
)

var bridge *huego.Bridge

func Init()  {
	bridge, _ = huego.Discover()
	bridge = bridge.Login(config.LoadedConfig.HueUser)
}

func BlinkAll()  {
	lights, _ := bridge.GetLights()
	for i := range lights {
		lights[i].Alert("lselect")
	}
}

func SetALlLightsColor(color string)  {
	SetAllLightsState(true)
	lights, _ := bridge.GetLights()

	c, _ := colorful.Hex(color)

	x, y, _ := c.Xyy()

	for i := range lights {
		lights[i].Xy([]float32{
			float32(x),
			float32(y),
		})
	}
}

func SetAllLightsState(on bool)  {
	lights, _ := bridge.GetLights()
	for i := range lights {
		if on {
			lights[i].On()
		} else {
			lights[i].Off()
		}
	}
}
