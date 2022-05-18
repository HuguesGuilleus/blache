// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2022, Hugues GUILLEUS
// All rights reserved.

package chunck

// One section, 16*16*16 block.
type Section struct {
	Y           uint8
	Palette     []PaletteItem
	BlockStates []int64
}

type PaletteItem struct {
	Name string
}

func (section *Section) decodeNBT(tagType byte, name string, r *reader) error {
	switch name {
	case "Y":
		if tagType != tagByte {
			return expectedTag(tagByte, tagType)
		}
		var err error
		section.Y, err = r.readByte()
		if err != nil {
			return err
		}

	case "Palette":
		if tagType != tagList {
			return expectedTag(tagList, tagType)
		}

		paletteLen, err := r.readListMeta(tagCompound)
		if err != nil {
			return err
		}

		section.Palette = make([]PaletteItem, paletteLen, paletteLen)
		for i := 0; i < paletteLen; i++ {
			if err := r.readCompound(section.Palette[i].decodeNBT); err != nil {
				return err
			}
		}

	case "BlockStates":
		if tagType != tagInt64Array {
			return expectedTag(tagInt64Array, tagType)
		}
		var err error
		section.BlockStates, err = r.readInt64Array()
		if err != nil {
			return err
		}

	default:
		return skipNode
	}
	return nil
}

func (item *PaletteItem) decodeNBT(tagType byte, name string, r *reader) (err error) {
	switch name {
	case "Name":
		if tagType != tagString {
			return expectedTag(tagString, tagType)
		}
		item.Name, err = r.readString()
		return
	default:
		return skipNode
	}
}
