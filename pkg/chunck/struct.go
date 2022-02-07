// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2022, Hugues GUILLEUS
// All rights reserved.

package chunck

// On column of 16*16*256 block = 16 section.
type Chunck struct {
	Level struct {
		Biomes     interface{}
		Sections   []Section
		Structures struct {
			Starts map[string]struct{}
		}
	}
}

// One section, 16*16*16 block.
type Section struct {
	Y       uint8
	Palette []struct {
		Name string
	}
	BlockStates []int64
}
