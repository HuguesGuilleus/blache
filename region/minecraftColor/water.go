// BSD 3-Clause License in LICENSE file at the project root.
// All rights reserved.

package minecraftColor

import (
	"image/color"
)

const HasWater uint8 = 1

var WaterPalette = color.Palette{
	color.RGBA{},                 // no water
	color.RGBA{64, 64, 255, 255}, // water
}
