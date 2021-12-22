// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2021, Hugues GUILLEUS
// All rights reserved.

// Unmarshal NBT data from []byte.
package nbt

import (
	"fmt"
	"reflect"
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

)

func Unmarshal(data []byte, destination interface{}) error {
	r := reader(data)
	return r.assignRoot(reflect.ValueOf(destination))
}

func (r *reader) unmarshal(value reflect.Value) error {
	tagType, err := r.readByte()
	if err != nil {
		return fmt.Errorf("Read tag type fail: %w", err)
	}

	switch tagType {
	case tagEnd:
		return fmt.Errorf("Unexpected tag End")
	case tagByte:
		return r.assignByte(value)
	case tagInt16:
		return unimplemented
	case tagInt32:
		return unimplemented
	case tagInt64:
		return unimplemented
	case tagFloat32:
		return unimplemented
	case tagFloat64:
		return unimplemented

	case tagString:
		return unimplemented

	case tagList:
		return unimplemented

	case tagCompound:
		return unimplemented

	case tagBytes:
		return unimplemented
	case tagInt32Array:
		return unimplemented
	case tagInt64Array:
		return unimplemented
	}
	return unimplemented
}

var unimplemented = fmt.Errorf("Not implemented")
