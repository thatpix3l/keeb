package rgb

import "image/color"

// Cycle RGBA struct, rainbow pattern
func Rainbow(rgba color.RGBA) color.RGBA {

	switch {

	case rgba.R < 255 && rgba.G == 0:
		rgba.R++

	case rgba.B > 0 && rgba.R == 255:
		rgba.B--

	case rgba.G < 255 && rgba.B == 0:
		rgba.G++

	case rgba.R > 0 && rgba.G == 255:
		rgba.R--

	case rgba.B < 255 && rgba.R == 0:
		rgba.B++

	case rgba.G > 0 && rgba.B == 255:
		rgba.G--

	}

	return rgba

}
