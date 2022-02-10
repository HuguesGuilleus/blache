// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2022, Hugues GUILLEUS
// All rights reserved.

package chunck

import (
	"fmt"
)

// Read bytes array value, introduced by TAG_Byte_Array.
func (r *reader) readBytesArray() ([]byte, error) {
	_, size, err := r.checkArrayLen(1)
	if err != nil {
		return nil, err
	}

	data := (*r)[:size]
	*r = (*r)[size:]
	return data, nil
}

// Read an []int32 array value, introduced by TAG_Int_Array.
func (r *reader) readInt32Array() ([]int32, error) {
	arrayLen, size, err := r.checkArrayLen(4)
	if err != nil {
		return nil, err
	}

	array := make([]int32, arrayLen, arrayLen)
	for i := range array {
		array[i] = int32((*r)[i*4+0])<<24 +
			int32((*r)[i*4+1])<<16 +
			int32((*r)[i*4+2])<<8 +
			int32((*r)[i*4+3])
	}

	*r = (*r)[size:]
	return array, nil
}

// Read an []int64 array value, introduced by TAG_Long_Array.
func (r *reader) readInt64Array() ([]int64, error) {
	arrayLen, size, err := r.checkArrayLen(8)
	if err != nil {
		return nil, err
	}

	array := make([]int64, arrayLen, arrayLen)
	for i := range array {
		array[i] = int64((*r)[i*8+0])<<56 +
			int64((*r)[i*8+1])<<48 +
			int64((*r)[i*8+2])<<40 +
			int64((*r)[i*8+3])<<32 +
			int64((*r)[i*8+4])<<24 +
			int64((*r)[i*8+5])<<16 +
			int64((*r)[i*8+6])<<8 +
			int64((*r)[i*8+7])
	}

	*r = (*r)[size:]
	return array, nil
}

// Read the len of this array, and check if the reader have the size.
func (r *reader) checkArrayLen(itemLen int) (arrayLen, size int, err error) {
	arrayLen, err = r.readLen()
	if err != nil {
		return
	}
	size = arrayLen * itemLen
	if len(*r) < size {
		err = fmt.Errorf("Array size (%d * %d byte) bigger NBT data (len:%d)", arrayLen, itemLen, len(*r))
		*r = nil
		return 0, 0, err
	}
	return
}
