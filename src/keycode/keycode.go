package keycode

import (
	"fmt"
	kbd "machine/usb/hid/keyboard"
)

type keycodeMap map[string]kbd.Keycode

var (
	keycodes = keycodeMap{

		"modifierctrl":       kbd.KeyModifierCtrl,
		"modifiershift":      kbd.KeyModifierShift,
		"modifieralt":        kbd.KeyModifierAlt,
		"modifiergui":        kbd.KeyModifierGUI,
		"modifierleftctrl":   kbd.KeyModifierLeftCtrl,
		"modifierleftshift":  kbd.KeyModifierLeftShift,
		"modifierleftalt":    kbd.KeyModifierLeftAlt,
		"modifierleftgui":    kbd.KeyModifierLeftGUI,
		"modifierrightctrl":  kbd.KeyModifierRightCtrl,
		"modifierrightshift": kbd.KeyModifierRightShift,
		"modifierrightalt":   kbd.KeyModifierRightAlt,
		"modifierrightgui":   kbd.KeyModifierRightGUI,

		"systempowerdown": kbd.KeySystemPowerDown,
		"systemsleep":     kbd.KeySystemSleep,
		"systemwakeup":    kbd.KeySystemWakeUp,

		"mediaplay":        kbd.KeyMediaPlay,
		"mediapause":       kbd.KeyMediaPause,
		"mediarecord":      kbd.KeyMediaRecord,
		"mediafastforward": kbd.KeyMediaFastForward,
		"mediarewind":      kbd.KeyMediaRewind,
		"medianexttrack":   kbd.KeyMediaNextTrack,
		"mediaprevtrack":   kbd.KeyMediaPrevTrack,
		"mediastop":        kbd.KeyMediaStop,
		"mediaeject":       kbd.KeyMediaEject,
		"mediarandomplay":  kbd.KeyMediaRandomPlay,
		"mediaplaypause":   kbd.KeyMediaPlayPause,
		"mediaplayskip":    kbd.KeyMediaPlaySkip,
		"mediamute":        kbd.KeyMediaMute,
		"mediavolumeinc":   kbd.KeyMediaVolumeInc,
		"mediavolumedec":   kbd.KeyMediaVolumeDec,

		"a":               kbd.KeyA,
		"b":               kbd.KeyB,
		"c":               kbd.KeyC,
		"d":               kbd.KeyD,
		"e":               kbd.KeyE,
		"f":               kbd.KeyF,
		"g":               kbd.KeyG,
		"h":               kbd.KeyH,
		"i":               kbd.KeyI,
		"j":               kbd.KeyJ,
		"k":               kbd.KeyK,
		"l":               kbd.KeyL,
		"m":               kbd.KeyM,
		"n":               kbd.KeyN,
		"o":               kbd.KeyO,
		"p":               kbd.KeyP,
		"q":               kbd.KeyQ,
		"r":               kbd.KeyR,
		"s":               kbd.KeyS,
		"t":               kbd.KeyT,
		"u":               kbd.KeyU,
		"v":               kbd.KeyV,
		"w":               kbd.KeyW,
		"x":               kbd.KeyX,
		"y":               kbd.KeyY,
		"z":               kbd.KeyZ,
		"1":               kbd.Key1,
		"2":               kbd.Key2,
		"3":               kbd.Key3,
		"4":               kbd.Key4,
		"5":               kbd.Key5,
		"6":               kbd.Key6,
		"7":               kbd.Key7,
		"8":               kbd.Key8,
		"9":               kbd.Key9,
		"0":               kbd.Key0,
		"enter":           kbd.KeyEnter,
		"esc":             kbd.KeyEsc,
		"backspace":       kbd.KeyBackspace,
		"tab":             kbd.KeyTab,
		"space":           kbd.KeySpace,
		"minus":           kbd.KeyMinus,
		"equal":           kbd.KeyEqual,
		"leftbrace":       kbd.KeyLeftBrace,
		"rightbrace":      kbd.KeyRightBrace,
		"backslash":       kbd.KeyBackslash,
		"nonusnum":        kbd.KeyNonUsNum,
		"semicolon":       kbd.KeySemicolon,
		"quote":           kbd.KeyQuote,
		"tilde":           kbd.KeyTilde,
		"comma":           kbd.KeyComma,
		"period":          kbd.KeyPeriod,
		"forwardslash":    kbd.KeySlash,
		"capslock":        kbd.KeyCapsLock,
		"f1":              kbd.KeyF1,
		"f2":              kbd.KeyF2,
		"f3":              kbd.KeyF3,
		"f4":              kbd.KeyF4,
		"f5":              kbd.KeyF5,
		"f6":              kbd.KeyF6,
		"f7":              kbd.KeyF7,
		"f8":              kbd.KeyF8,
		"f9":              kbd.KeyF9,
		"f10":             kbd.KeyF10,
		"f11":             kbd.KeyF11,
		"f12":             kbd.KeyF12,
		"printscreen":     kbd.KeyPrintscreen,
		"scrolllock":      kbd.KeyScrollLock,
		"pause":           kbd.KeyPause,
		"insert":          kbd.KeyInsert,
		"home":            kbd.KeyHome,
		"pageup":          kbd.KeyPageUp,
		"delete":          kbd.KeyDelete,
		"end":             kbd.KeyEnd,
		"pagedown":        kbd.KeyPageDown,
		"right":           kbd.KeyRight,
		"left":            kbd.KeyLeft,
		"down":            kbd.KeyDown,
		"up":              kbd.KeyUp,
		"numlock":         kbd.KeyNumLock,
		"padforwardslash": kbd.KeypadSlash,
		"padasterisk":     kbd.KeypadAsterisk,
		"padminus":        kbd.KeypadMinus,
		"padplus":         kbd.KeypadPlus,
		"padenter":        kbd.KeypadEnter,
		"pad1":            kbd.Keypad1,
		"pad2":            kbd.Keypad2,
		"pad3":            kbd.Keypad3,
		"pad4":            kbd.Keypad4,
		"pad5":            kbd.Keypad5,
		"pad6":            kbd.Keypad6,
		"pad7":            kbd.Keypad7,
		"pad8":            kbd.Keypad8,
		"pad9":            kbd.Keypad9,
		"pad0":            kbd.Keypad0,
		"padperiod":       kbd.KeypadPeriod,
		"nonusbs":         kbd.KeyNonUSBS,
		"menu":            kbd.KeyMenu,
		"f13":             kbd.KeyF13,
		"f14":             kbd.KeyF14,
		"f15":             kbd.KeyF15,
		"f16":             kbd.KeyF16,
		"f17":             kbd.KeyF17,
		"f18":             kbd.KeyF18,
		"f19":             kbd.KeyF19,
		"f20":             kbd.KeyF20,
		"f21":             kbd.KeyF21,
		"f22":             kbd.KeyF22,
		"f23":             kbd.KeyF23,
		"f24":             kbd.KeyF24,

		"uparrow":    kbd.KeyUpArrow,
		"downarrow":  kbd.KeyDownArrow,
		"leftarrow":  kbd.KeyLeftArrow,
		"rightarrow": kbd.KeyRightArrow,
		"return":     kbd.KeyReturn,
		"leftctrl":   kbd.KeyLeftCtrl,
		"leftshift":  kbd.KeyLeftShift,
		"leftalt":    kbd.KeyLeftAlt,
		"leftgui":    kbd.KeyLeftGUI,
		"rightctrl":  kbd.KeyRightCtrl,
		"rightshift": kbd.KeyRightShift,
		"rightalt":   kbd.KeyRightAlt,
		"rightgui":   kbd.KeyRightGUI,
	}
)

// Given a lower-case name, find the corresponding keycode
func Get(name string) (kbd.Keycode, error) {

	code, ok := keycodes[name]
	if !ok {
		return kbd.Key0, fmt.Errorf("failure to get keycode: non-existent name \"%s\"", name)
	}

	return code, nil
}
