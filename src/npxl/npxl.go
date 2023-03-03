package npxl

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ws2812"
)

var (
	Onboard = new(machine.GPIO25)
)

type neopixel struct {
	ws2812.Device
}

func (n *neopixel) Orange() {
	n.WriteColors([]color.RGBA{{R: 255, G: 85, B: 0}})
}

func (n *neopixel) Blue() {
	n.WriteColors([]color.RGBA{{R: 0, G: 0, B: 255}})
}

// Turn a pin into an addressable WS2812 device
func new(pin machine.Pin) neopixel {
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	device := ws2812.New(pin)
	return neopixel{device}
}
