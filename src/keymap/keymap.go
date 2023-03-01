package keymap

import (
	"image/color"
	"sync"

	"github.com/thatpix3l/keeb/src/board"
)

var (
	wgForeach sync.WaitGroup
)

type ledEffect string

const (
	Breathing ledEffect = "breathing"
	Static    ledEffect = "static"
)

type led struct {
	effect ledEffect
	color  color.RGBA
}

// 2D map of rgb pointers
type layerLEDs map[int]map[int]*led

// Key properties for the currently active layer
type LayerKey struct {
	TempLayer         string   // Layer name to temporarily switch to when pressing and holding key
	Sequence          []string // Sequence of keys to press in given order, and release in reversed order
	PreviouslyPressed bool     // If the key was previously pressed
}

type layerKeys map[int]map[int]*LayerKey

// Layer containing one matrix of keys and LEDs.
type layer struct {
	Keys layerKeys
	Leds layerLEDs
}

// 1D Map of layer pointers
type layers map[string]*layer

type defaultLayerName string

// Root keymap. Do NOT use directly to instantiate.
type Mapping struct {
	layers           layers
	DefaultLayerName defaultLayerName
	CurrentLayerName string
}

// Proper way of creating a new mapping.
func New() Mapping {
	newMapping := Mapping{
		layers:           layers{},
		DefaultLayerName: "base",
	}
	newMapping.CurrentLayerName = string(newMapping.DefaultLayerName)

	return newMapping
}

// Access a keyboard layer. If it doesn't exist, create one.
func (m *Mapping) Layer(name string) *layer {
	if m.layers[name] == nil {
		m.layers[name] = &layer{}
	}
	return m.layers[name]
}

// Run callback for each physical key row and column in layer
func (l *layer) ForEach(callback func(row int, col int, key *LayerKey)) {

	// If no map of key rows, create one
	if l.Keys == nil {
		l.Keys = map[int]map[int]*LayerKey{}
	}

	// For each row...
	for rowIdx := 0; rowIdx < len(board.Config.RowPins); rowIdx++ {

		// If no map of key columns in current row, create one
		if l.Keys[rowIdx] == nil {
			l.Keys[rowIdx] = map[int]*LayerKey{}
		}

		// For each column...
		for colIdx := 0; colIdx < board.Config.MaxKeysPerRow[rowIdx]; colIdx++ {

			// If no key in current row and column, create one
			if l.Keys[rowIdx][colIdx] == nil {
				l.Keys[rowIdx][colIdx] = &LayerKey{}
			}

			callback(rowIdx, colIdx, l.Keys[rowIdx][colIdx])

		}

	}

}
