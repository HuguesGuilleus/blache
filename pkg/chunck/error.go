// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2022, Hugues GUILLEUS
// All rights reserved.

package chunck

import (
	"fmt"
)

func expectedTagCompound(tagType byte) error {
	return expectedTagType("TAG_Compound", tagType)
}
func expectedTagList(tagType byte) error {
	return expectedTagType("TAG_List", tagType)
}
func expectedTagType(expected string, recevied byte) error {
	receviedName := ""
	switch recevied {
	case tagEnd:
		receviedName = "TAG_End"
	case tagByte:
		receviedName = "TAG_Byte"
	case tagInt16:
		receviedName = "TAG_Short"
	case tagInt32:
		receviedName = "TAG_Int"
	case tagInt64:
		receviedName = "TAG_Long"
	case tagFloat32:
		receviedName = "TAG_Float"
	case tagFloat64:
		receviedName = "TAG_Double"
	case tagBytes:
		receviedName = "TAG_Byte_Array"
	case tagString:
		receviedName = "TAG_String"
	case tagList:
		receviedName = "TAG_List"
	case tagCompound:
		receviedName = "TAG_Compound"
	case tagInt32Array:
		receviedName = "TAG_Int_Array"
	case tagInt64Array:
		receviedName = "TAG_Long_Array"
	}
	return fmt.Errorf("Expected a %s type, found: %s", expected, receviedName)
}
