// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2022, Hugues GUILLEUS
// All rights reserved.

package chunck

import (
	"fmt"
)

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

func (c *Chunck) DecodeNBT(data []byte) error {
	r := reader(data)
	for {
		tagType, name, err := r.readTagMeta()
		if err != nil {
			return fmt.Errorf("Decode Chunck fail: %w", err)
		} else if tagType == tagEnd {
			break
		}
		switch name {
		case "":
			if tagType != tagCompound {
				return expectedTagCompound(tagType)
			}
			for {
				tagType, name, err := r.readTagMeta()
				if err != nil {
					return fmt.Errorf("Decode Chunck fail: %w", err)
				} else if tagType == tagEnd {
					break
				}

				switch name {
				case "Level":

					if tagType != tagCompound {
						return expectedTagCompound(tagType)
					}
					return c.decodeNBT(&r)
				default:
					if err := r.skip(tagType); err != nil {
						return fmt.Errorf("Decode Chunck.%q: %w", name, err)
					}
				}
			}
		default:
			if err := r.skip(tagType); err != nil {
				return fmt.Errorf("Decode Chunck.%q: %w", name, err)
			}
		}
	}

	// The draw function can use a empty chunck.
	return nil
}

func (c *Chunck) decodeNBT(r *reader) error {
	for {
		tagType, name, err := r.readTagMeta()
		if err != nil {
			return fmt.Errorf("Read Chunck.Level fail: %w", err)
		} else if tagType == tagEnd {
			break
		}

		switch name {
		case "Sections":
			if tagType != tagList {
				return expectedTagList(tagType)
			}
			if sectionType, err := r.readByte(); err != nil {
				return err
			} else if sectionType != tagCompound {
				return fmt.Errorf("Read Chunck.Level.Sections: %w", expectedTagCompound(tagType))
			}
			listLen, err := r.readLen()
			if err != nil {
				return err
			}
			c.Level.Sections = make([]Section, listLen, listLen)
			for i := 0; i < listLen; i++ {
				if err := r.readCompound(c.Level.Sections[i].decodeNBT); err != nil {
					return err
				}
			}

		case "Structures":
			if tagType != tagCompound {
				return fmt.Errorf("Decode Chunck.Level.Structures: %w", expectedTagCompound(tagType))
			}
			if err := c.nbtLevelStructure(r); err != nil {
				return fmt.Errorf("Decode Chunck.Level.Structures: %w", err)
			}

		default:
			if err := r.skip(tagType); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Chunck) nbtLevelStructure(r *reader) error {
	for {
		tagType, name, err := r.readTagMeta()
		if err != nil {
			return err
		} else if tagType == tagEnd {
			break
		}

		switch name {
		case "Starts":
			if tagType != tagCompound {
				return fmt.Errorf("Read Chunck.Level.Structures[].Strats: %w", expectedTagCompound(tagType))
			}
			c.Level.Structures.Starts = make(map[string]struct{})
			for {
				tagType, name, err := r.readTagMeta()
				if err != nil {
					return err
				} else if tagType == tagEnd {
					break
				}
				c.Level.Structures.Starts[name] = struct{}{}
				return r.skip(tagType)
			}
		default:
			if err := r.skip(tagType); err != nil {
				return fmt.Errorf("Skip Level.Structure.%q fail: %w", name, err)
			}
		}
	}
	return nil
}
