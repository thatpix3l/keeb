// Board config for mapping between pin rows/columns and physical keys.
package board

import (
	m "machine"
)

type pins []m.Pin

// Run callback for each pin
func (p pins) ForEach(callback func(m.Pin)) {
	for _, pin := range p {
		callback(pin)
	}
}

type config struct {
	RowPins       pins
	ColPins       pins
	MaxKeysPerRow []int
}

var (
	Config = config{
		RowPins:       []m.Pin{m.GPIO9, m.GPIO21, m.GPIO23, m.GPIO20, m.GPIO22},
		ColPins:       []m.Pin{m.GPIO2, m.GPIO3, m.GPIO4, m.GPIO5, m.GPIO6, m.GPIO7, m.GPIO8},
		MaxKeysPerRow: []int{14, 14, 13, 12, 8},
	}
)
