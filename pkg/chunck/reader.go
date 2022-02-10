// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2022, Hugues GUILLEUS
// All rights reserved.

package chunck

import (
	"errors"
	"fmt"
	"io"
	"math"
)

const (
	tagEnd        byte = iota // TAG_End
	tagByte                   // TAG_Byte
	tagInt16                  // TAG_Short
	tagInt32                  // TAG_Int
	tagInt64                  // TAG_Long
	tagFloat32                // TAG_Float
	tagFloat64                // TAG_Double
	tagBytes                  // TAG_Byte_Array
	tagString                 // TAG_String
	tagList                   // TAG_List
	tagCompound               // TAG_Compound
	tagInt32Array             // TAG_Int_Array
	tagInt64Array             // TAG_Long_Array

	// All valid tag are less than this constants.
	tagBigger = tagInt64Array
)

var (
	// Fake error used to ask to skip this node.
	skipNode = errors.New("Skip this node")
	// Fake error to exit the decoding.
	exitWalk = errors.New("Exit node walk")
)

// The reader
type reader []byte

// Read a NBT tree.
func (r *reader) readTree(saver func(tagType byte, name string, r *reader) error) error {
	return r.readCompound(func(tagType byte, _ string, r *reader) error {
		if tagType != tagCompound {
			return expectedTag(tagCompound, tagType)
		}
		if err := r.readCompound(saver); err != nil {
			return err
		}
		return exitWalk
	})
}

// Read a NBT Compound, and to each
func (r *reader) readCompound(saver func(tagType byte, name string, r *reader) error) error {
	for {
		tagType, err := r.readByte()
		if err != nil {
			return fmt.Errorf("Read tag type: %w", err)
		} else if tagType > tagBigger {
			return invalidTag(tagType)
		} else if tagType == tagEnd {
			return nil
		}

		name, err := r.readString()
		if err != nil {
			return fmt.Errorf("Read tag name: %w", err)
		}

		switch err := saver(tagType, name, r); err {
		case nil:
			// ok
		case exitWalk:
			return nil
		case skipNode:
			if err := r.skip(tagType); err != nil {
				return fmt.Errorf("Skip %q: %w", name, err)
			}
		default:
			return fmt.Errorf("Decode %q: %w", name, err)
		}
	}
}

// Skip the value mabe nested of the tag.
func (r *reader) skip(tagType byte) error {
	switch tagType {
	case tagByte:
		r.skipBytes(1)
	case tagInt16:
		r.skipBytes(2)
	case tagInt32:
		r.skipBytes(4)
	case tagInt64:
		r.skipBytes(8)
	case tagFloat32:
		r.skipBytes(4)
	case tagFloat64:
		r.skipBytes(8)

	case tagString:
		_, err := r.readString()
		return err

	case tagBytes:
		return r.skipArray(1)
	case tagInt32Array:
		return r.skipArray(4)
	case tagInt64Array:
		return r.skipArray(8)

	case tagList:
		tagType, err := r.readByte()
		if err != nil {
			return fmt.Errorf("Read tag type (int8) fail: %w", err)
		}
		listLen, err := r.readLen()
		if err != nil {
			return err
		}
		for i := 0; i < listLen; i++ {
			if err := r.skip(tagType); err != nil {
				return err
			}
		}
	case tagCompound:
		return r.readCompound(func(_ byte, _ string, _ *reader) error {
			return skipNode
		})
	}

	return nil
}

// Skip the array data.
func (r *reader) skipArray(elemSize int) error {
	len, err := r.readLen()
	if err != nil {
		return fmt.Errorf("Skip array: %w", err)
	}
	r.skipBytes(len * elemSize)
	return nil
}

// Skip the next bytes in the stram reader.
func (r *reader) skipBytes(skipBytes int) {
	if skipBytes < len(*r) {
		*r = (*r)[skipBytes:]
	} else {
		*r = nil
	}
}

// Check the list item type and read the lenght of the list.
//
// If the type if TAG_End and the list is empty, no return error.
func (r *reader) readListMeta(expectedTagType byte) (listLen int, err error) {
	var tagType byte
	tagType, err = r.readByte()
	if err != nil {
		return 0, fmt.Errorf("Read list item tagType: %w", err)
	} else if tagType > tagBigger {
		return 0, invalidTag(tagType)
	}

	listLen, err = r.readLen()
	if err != nil {
		return 0, err
	}

	if tagType == tagEnd && listLen == 0 {
		return 0, nil
	} else if tagType != expectedTagType {
		return 0, fmt.Errorf("list item type %w", expectedTag(expectedTagType, tagType))
	}

	return
}

// Read playload of a NBT array of bytes, int32 or int64.
func (r *reader) readArray(itemLen int) (arrayLen int, data []byte, err error) {
	arrayLen, err = r.readLen()
	if err != nil {
		return 0, nil, err
	}
	size := arrayLen * itemLen
	if len(*r) < size {
		return 0, nil, fmt.Errorf("Array size (%d * %d byte) bigger NBT data (len:%d)", arrayLen, itemLen, len(*r))
	}
	data = (*r)[:size]
	*r = (*r)[size:]
	return arrayLen, data, nil
}

// Read the len of an Array or a List.
func (r *reader) readLen() (int, error) {
	len32, err := r.readInt32()
	if err != nil {
		return 0, fmt.Errorf("Read len (int32) fail: %w", err)
	}
	return int(len32), nil
}

/* Primitive value sreader */

// Read a primitive byte.
func (r *reader) readByte() (b byte, err error) {
	if len(*r) == 0 {
		return 0, fmt.Errorf("Read byte fail because %w", io.ErrUnexpectedEOF)
	}
	b = (*r)[0]
	*r = (*r)[1:]
	return
}

// Read a primitive int16.
func (r *reader) readInt16() (i int16, err error) {
	if len(*r) < 2 {
		return 0, fmt.Errorf("Read int16 fail because %w", io.ErrUnexpectedEOF)
	}

	i = int16((*r)[0])<<8 +
		int16((*r)[1])

	*r = (*r)[2:]
	return
}

// Read a primitive int32.
func (r *reader) readInt32() (i int32, err error) {
	if len(*r) < 4 {
		return 0, fmt.Errorf("Read int32 fail because %w", io.ErrUnexpectedEOF)
	}

	i = int32((*r)[0])<<24 +
		int32((*r)[1])<<16 +
		int32((*r)[2])<<8 +
		int32((*r)[3])

	*r = (*r)[4:]
	return
}

// Read a primitive int64.
func (r *reader) readInt64() (i int64, err error) {
	if len(*r) < 8 {
		return 0, fmt.Errorf("Read int64 fail because %w", io.ErrUnexpectedEOF)
	}

	i = int64((*r)[0])<<56 +
		int64((*r)[1])<<48 +
		int64((*r)[2])<<40 +
		int64((*r)[3])<<32 +
		int64((*r)[4])<<24 +
		int64((*r)[5])<<16 +
		int64((*r)[6])<<8 +
		int64((*r)[7])

	*r = (*r)[8:]
	return
}

// Read a primitive float32.
func (r *reader) readFloat32() (float32, error) {
	i, err := r.readInt32()
	if err != nil {
		return 0, fmt.Errorf("Read float32 fail because %w", io.ErrUnexpectedEOF)
	}
	return math.Float32frombits(uint32(i)), nil
}

// Read a primitive float64
func (r *reader) readFloat64() (float64, error) {
	i, err := r.readInt64()
	if err != nil {
		return 0, fmt.Errorf("Read float64 fail because %w", io.ErrUnexpectedEOF)
	}
	return math.Float64frombits(uint64(i)), nil
}

// Read a NBT primitive string (len + content).
func (r *reader) readString() (s string, err error) {
	var strLen16 int16
	strLen16, err = r.readInt16()
	strLen := int(strLen16)
	if err != nil {
		return "", fmt.Errorf("Read string len (int16) fail because %w", io.ErrUnexpectedEOF)
	}

	if len(*r) < strLen {
		return "", fmt.Errorf("Read string content fail because %w", io.ErrUnexpectedEOF)
	}
	s = string((*r)[:strLen])
	*r = (*r)[strLen:]
	return
}
