// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2021, Hugues GUILLEUS
// All rights reserved.

package nbt

import (
	_ "embed"
	"fmt"
	"io"
	"testing"
)

const tab string = "   "

//go:embed bigtest.nbt
var bigtest []byte

func TestChunck(t *testing.T) {
	r := reader(bigtest)
	t.Error(r.walk())
}

// Read all the NBT tree.
func (r *reader) walk() error {
	return r.walkNamedTag("")
}

// Read one item of a map.
func (r *reader) walkNamedTag(prefix string) error {
	tagType, err := r.readByte()
	if err != nil {
		return fmt.Errorf("Read tag type fail: %w", err)
	} else if tagType == tagEnd {
		return io.EOF
	}

	name, err := r.readString()
	if err != nil {
		return fmt.Errorf("Fail to read the name: %w", err)
	}

	fmt.Printf("%s%q: ", prefix, name)

	return r.walkContent(prefix, tagType)
}

// Read the content of the NBT node, with tagType already known.
func (r *reader) walkContent(prefix string, tagType byte) (err error) {
	p := func(v interface{}, err error) error {
		fmt.Printf("%T:%#v 0x%x\n", v, v, v)
		return err
	}

	switch tagType {
	// primary type
	case tagByte:
		p(r.readByte())
	case tagInt16:
		p(r.readInt16())
	case tagInt32:
		p(r.readInt32())
	case tagInt64:
		p(r.readInt64())
	case tagFloat32:
		p(r.readFloat32())
	case tagFloat64:
		p(r.readFloat64())
	case tagString:
		p(r.readString())

	// One type array
	case tagBytes:
		err = r.skipArray(1)
		fmt.Println("[]byte")
	case tagInt32Array:
		err = r.skipArray(4)
		fmt.Println("[]int32")
	case tagInt64Array:
		err = r.skipArray(8)
		fmt.Println("[]int64")

	// Complex type
	case tagCompound:
		fmt.Printf("@map {\n")
		defer fmt.Printf("%s}\n", prefix)
		for {
			err := r.walkNamedTag(prefix + tab)
			switch err {
			case nil:
				continue
			case io.EOF:
				return nil
			default:
				return err
			}
		}

	case tagList:
		itemType, err := r.readByte()
		if err != nil {
			return fmt.Errorf("Read type of item fail: %w", err)
		}
		_size, err := r.readInt32()
		if err != nil {
			return fmt.Errorf("Read of array size fail: %w", err)
		}
		size := int(_size)

		fmt.Printf("[%d]item [\n", size)
		defer fmt.Printf("%s]\n", prefix)

		for i := 0; i < size; i++ {
			fmt.Printf("%s[%d]: ", prefix+tab, i)
			if err = r.walkContent(prefix+tab, itemType); err != nil {
				return fmt.Errorf("Read the item %d fail: %w", i, err)
			}
		}

	// Fail type
	case tagEnd:
		return fmt.Errorf("Unexpected Tag_End")
	default:
		return fmt.Errorf("Unknown this tag: %d", tagType)
	}

	return err
}

func (r *reader) skipArray(sizeof int) error {
	l, err := r.readInt32()
	if err != nil {
		return fmt.Errorf("Fail to read the size of the arrray: %w", err)
	}
	size := int(l) * sizeof
	if len(*r) < size {
		return fmt.Errorf("Read array item.sizeof:%d, len:%d, reader.len:%d", sizeof, l, len(*r))
	}
	*r = (*r)[size:]

	return nil
}
