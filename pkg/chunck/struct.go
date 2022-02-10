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
		c.Level.Structures.Starts = make(map[string]struct{})
		return r.readCompound(func(_ byte, name string, _ *reader) error {
			c.Level.Structures.Starts[name] = struct{}{}
			return skipNode
		})
	default:
		return skipNode
	}
}
