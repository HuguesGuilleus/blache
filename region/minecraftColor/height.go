// BSD 3-Clause License in LICENSE file at the project root.
// All rights reserved.

package minecraftColor

import (
	"image/color"
)

// Standard palette for Height, color.Gray inside.
var HeightPalette color.Palette = func() (palette color.Palette) {
	palette = make(color.Palette, 256)
	for i := range palette {
		palette[i] = color.Gray{uint8(i)}
	}
	return
}()
