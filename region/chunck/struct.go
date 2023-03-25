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

func (chunck *Chunck) DecodeNBT(data []byte) error {
	r := reader(data)
	return r.readTree(chunck.decodeChunck)
}

func (chunck *Chunck) decodeChunck(tagType byte, name string, r *reader) error {
	switch name {
	case "Level":
		if tagType != tagCompound {
			return expectedTag(tagCompound, tagType)
		}
		if err := r.readCompound(chunck.decodeLevel); err != nil {
			return err
		}
		return exitWalk
	default:
		return skipNode
	}
}

func (chunck *Chunck) decodeLevel(tagType byte, name string, r *reader) error {
	switch name {
	case "Biomes":
		var err error
		switch tagType {
		case tagBytes:
			chunck.Level.Biomes, err = r.readBytesArray()
		case tagInt32Array:
			chunck.Level.Biomes, err = r.readInt32Array()
		default:
			return skipNode
		}
		return err
	case "Sections":
		if tagType != tagList {
			return expectedTag(tagList, tagType)
		}
		listLen, err := r.readListMeta(tagCompound)
		if err != nil {
			return err
		}
		chunck.Level.Sections = make([]Section, listLen, listLen)
		for i := range chunck.Level.Sections {
			if err := r.readCompound(chunck.Level.Sections[i].decodeNBT); err != nil {
				return err
			}
		}

	case "Structures":
		if tagType != tagCompound {
			return expectedTag(tagCompound, tagType)
		} else if err := r.readCompound(chunck.nbtLevelStructure); err != nil {
			return err
		}

	default:
		return skipNode
	}
	return nil
}

func (c *Chunck) nbtLevelStructure(tagType byte, name string, r *reader) error {
	switch name {
	case "Starts":
		if tagType != tagCompound {
			return expectedTag(tagCompound, tagType)
		}
		starts := make(map[string]struct{})
		c.Level.Structures.Starts = starts
		return r.readCompound(func(tagType byte, structureName string, r *reader) error {
			if tagType != tagCompound {
				return skipNode
			}
			r.readCompound(func(tagType byte, name string, r *reader) error {
				if tagType != tagString || name != "id" {
					return skipNode
				} else if s, err := r.readString(); err != nil {
					return err
				} else if s == "INVALID" {
					return nil
				}
				starts[structureName] = struct{}{}
				return nil
			})
			return nil
		})
	default:
		return skipNode
	}
}
