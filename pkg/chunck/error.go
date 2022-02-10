// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2022, Hugues GUILLEUS
// All rights reserved.

package chunck

import (
	"fmt"
)

func expectedTagCompound(tag byte) error {
	return expectedTag(tagCompound, tag)
}
func expectedTagList(tag byte) error {
	return expectedTag(tagList, tag)
}
func expectedTagType(expected string, received byte) error {
	return fmt.Errorf("Expected a %s type, found: %s", expected, tag2string(received))
}

// Create a fmt error
func expectedTag(expected, received byte) error {
	return fmt.Errorf("Expected a %s but found %s", tag2string(expected), tag2string(received))
}

// Return the tag in string with the based format.
func tag2string(tag byte) string {
	switch tag {
	case tagEnd:
		return "TAG_End"
	case tagByte:
		return "TAG_Byte"
	case tagInt16:
		return "TAG_Short"
	case tagInt32:
		return "TAG_Int"
	case tagInt64:
		return "TAG_Long"
	case tagFloat32:
		return "TAG_Float"
	case tagFloat64:
		return "TAG_Double"
	case tagBytes:
		return "TAG_Byte_Array"
	case tagString:
		return "TAG_String"
	case tagList:
		return "TAG_List"
	case tagCompound:
		return "TAG_Compound"
	case tagInt32Array:
		return "TAG_Int_Array"
	case tagInt64Array:
		return "TAG_Long_Array"
	default:
		return fmt.Sprintf("Unknown tag type: %d", tag)
	}
}
