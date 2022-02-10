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

		itemType, paletteLen, err := r.readListMeta()
		if err != nil {
			return err
		} else if itemType != tagCompound {
			return expectedTag(tagCompound, tagType)
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
		l, bytes, err := r.readArray(8)
		if err != nil {
			return err
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
