// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2021, Hugues GUILLEUS
// All rights reserved.

package minecraftColor

import (
	"image/color"
)

var HeightPalette color.Palette = func() (palette color.Palette) {
	palette = make(color.Palette, 256)
	for i := range palette {
		palette[i] = color.Gray{uint8(i)}
	}
	return
}()
