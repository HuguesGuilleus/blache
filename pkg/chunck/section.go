// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2022, Hugues GUILLEUS
// All rights reserved.

package chunck

import (
	"fmt"
)

// One section, 16*16*16 block.
type Section struct {
	Y       uint8
	Palette []struct {
		Name string
	}
	BlockStates []int64
}

func (section *Section) nbt(r *reader) error {
	for {
		tagType, name, err := r.readTagMeta()
		if err != nil {
			return err
		} else if tagType == tagEnd {
			break
		}

		switch name {
		case "Y":
			if tagType != tagByte {
				return fmt.Errorf("Decode Chunck.Level.Sections[].Y: %w", err)
			}
			section.Y, err = r.readByte()
			if err != nil {
				return err
			}

		case "Palette":
			if tagType != tagList {
				return expectedTagList(tagType)
			}
			if itemType, err := r.readByte(); err != nil {
				return err
			} else if itemType != tagCompound {
				return expectedTagCompound(itemType)
			}
			paletteLen, err := r.readLen()
			if err != nil {
				return err
			}
			section.Palette = make([]struct{ Name string }, paletteLen, paletteLen)
			for i := 0; i < paletteLen; i++ {
				for {
					tagType, name, err := r.readTagMeta()
					if err != nil {
						return err
					} else if tagType == tagEnd {
						break
					}

					switch name {
					case "Name":
						if tagType != tagString {
							return expectedTagType("TAG_String", tagType)
						}
						section.Palette[i].Name, err = r.readString()
						if err != nil {
							return err
						}
					default:
						r.skip(tagType)
					}
				}
			}

		case "BlockStates":
			if tagType != tagInt64Array {
				return expectedTagType("TAG_Long_Array", tagType)
			}
			l, bytes, err := r.readArray(8)
			if err != nil {
				return fmt.Errorf("Read Chuck.Level.Sections[].BlockStates: %w", err)
			}
			section.BlockStates = make([]int64, l, l)
			for i := range section.BlockStates {
				section.BlockStates[i] = int64(bytes[i*8+0])<<56 +
					int64(bytes[i*8+1])<<48 +
					int64(bytes[i*8+2])<<40 +
					int64(bytes[i*8+3])<<32 +
					int64(bytes[i*8+4])<<24 +
					int64(bytes[i*8+5])<<16 +
					int64(bytes[i*8+6])<<8 +
					int64(bytes[i*8+7])
			}

		default:
			r.skip(tagType)
		}
	}
	return nil
}
