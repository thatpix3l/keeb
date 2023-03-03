package main

import (
	"image/color"
	"log"
	"machine"
	"time"

	kbd "machine/usb/hid/keyboard"

	"github.com/thatpix3l/keeb/src/board"
	"github.com/thatpix3l/keeb/src/keycode"
	"github.com/thatpix3l/keeb/src/keymap"
	"github.com/thatpix3l/keeb/src/rgb"
	"tinygo.org/x/drivers/ws2812"
)

// Get access to the onboard WS2812 RGB LED
func newNeopixel(pin machine.Pin) *ws2812.Device {
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	newPin := ws2812.New(pin)
	return &newPin
}

var npxl = newNeopixel(machine.GPIO25)

// Change onboard LED color to blue
func ledBlue() {
	npxl.WriteColors([]color.RGBA{{R: 0, G: 0, B: 255}})
}

// Change onboard LED color to orange
func ledOrange() {
	npxl.WriteColors([]color.RGBA{{R: 255, G: 85, B: 0}})
}

// Create a callback that accepts a pin and configures it to the given mode
func pinMode(mode machine.PinMode) func(machine.Pin) {
	return func(p machine.Pin) {
		p.Configure(machine.PinConfig{Mode: mode})
	}
}

// Callback func type for key press or release
type pinKeyCallback func(rowKeyIdx int, colKeyIdx int)

// Struct with fields containing events that fire when a key is pressed and released
type scanner struct {
	onPress   pinKeyCallback // Key callback when pressed
	onRelease pinKeyCallback // Key callback when released
}

// Yes, this is a thing. No, I won't remove it.
func isEven(num int) bool {
	return num%2 == 0
}

// Yes, this is a thing. No, I won't remove it.
func isOdd(num int) bool {
	return num%2 != 0
}

// Algorithm for board pins between column-to-row and row-to-column.
func japaneseDuplex(matrix [][]bool) [][]bool {

	for keyRowIdx, keyRow := range matrix {

		rowPin := board.Config.RowPins[keyRowIdx]

		for keyColIdx := range keyRow {

			pinColIdx := int(float64(keyColIdx) / 2)
			colPin := board.Config.ColPins[pinColIdx]

			var outPin machine.Pin
			var inPin machine.Pin

			if isEven(keyColIdx) {
				outPin = colPin
				inPin = rowPin

			} else {
				outPin = rowPin
				inPin = colPin
			}

			outPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
			outPin.High()

			matrix[keyRowIdx][keyColIdx] = inPin.Get()

			outPin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

		}

	}

	return matrix

}

// Scan, in one cycle, for pressed and released keys, using a given algorithm
func (s *scanner) scanCycle(matrix [][]bool, algo func(matrix [][]bool) [][]bool) {

	board.Config.RowPins.ForEach(pinMode(machine.PinInputPulldown))
	board.Config.ColPins.ForEach(pinMode(machine.PinInputPulldown))

	algo(matrix)

	// After all the algorithms modified the matrix, check which are pressing right now
	for row, colPressing := range matrix {
		for col, pressing := range colPressing {

			if pressing {
				s.onPress(row, col)

			} else {
				s.onRelease(row, col)
			}

		}
	}

}

// Create a new key status matrix. It is based off of the amount of board pin rows/columns and how many keys there are per row.
func initKeyMatrix() [][]bool {

	newKeyMatrix := [][]bool{}

	for row := 0; row < len(board.Config.RowPins); row++ {
		newKeyMatrix = append(newKeyMatrix, []bool{})

		for col := 0; col < board.Config.MaxKeysPerRow[row]; col++ {
			newKeyMatrix[row] = append(newKeyMatrix[row], false)
		}
	}

	return newKeyMatrix
}

// Scan, infinitely, for pressed and released keys
func (s scanner) scanLoop() {
	keyMatrix := initKeyMatrix()
	for {
		s.scanCycle(keyMatrix, japaneseDuplex)
	}
}

var ledColor = color.RGBA{0, 0, 0, 0}

// Cycle onboard LED color in rainbow pattern
func ledRainbow() {
	npxl.WriteColors([]color.RGBA{ledColor})
	ledColor = rgb.Rainbow(ledColor)
	time.Sleep(time.Second / 255)
}

func bootupLightAnimation() {
	past := time.Now()
	for {
		ledRainbow()
		present := time.Now()

		difference := present.Sub(past)

		if difference.Seconds() > 3 {
			break
		}

	}
}

func main() {

	mainKeymap := keymap.New()

	// For each key in the base layer...
	mainKeymap.Layer("base").ForEach(func(row, col int, key *keymap.LayerKey) {

		// If a row or a row's column has not been configured, skip it
		if keymap.DefaultSequences[row] == nil || keymap.DefaultSequences[row][col] == nil {
			return
		}

		// Assign the current key sequence to the current key
		key.Sequence = keymap.DefaultSequences[row][col]

	})

	kbdMain := kbd.New()

	// Lightup RGB barf animation
	bootupLightAnimation()

	// New key scanner
	scnr := scanner{}

	// Pretty basic press and release callbacks. For now, only key sequence press and release, in appropriate order, are implemented
	scnr.onPress = func(row int, col int) {

		key := mainKeymap.Layer(mainKeymap.CurrentLayerName).Keys[row][col]
		if key.PreviouslyPressed {
			return
		}

		// At this point, the user IS pressing the key and was NOT pressed before now

		ledBlue()

		for _, keyName := range key.Sequence {

			kcode, err := keycode.Get(keyName)

			// TODO: handle temp layer switching when key is pressed

			if err != nil {
				log.Println(err)
				return
			}

			kbdMain.Down(kcode)
		}

		key.PreviouslyPressed = true

	}

	scnr.onRelease = func(row int, col int) {

		key := mainKeymap.Layer(mainKeymap.CurrentLayerName).Keys[row][col]
		if !key.PreviouslyPressed {
			return
		}

		// At this point, the user is NOT pressing the key and WAS pressed before now

		ledOrange()

		for seqIdx := range key.Sequence {

			keyName := key.Sequence[len(key.Sequence)-1-seqIdx]
			kcode, err := keycode.Get(keyName)

			// TODO: handle temp layer switching when key is pressed

			if err != nil {
				log.Println(err)
				return
			}

			kbdMain.Up(kcode)

		}

		key.PreviouslyPressed = false

	}

	// Blocking scanner loop
	scnr.scanLoop()

}
