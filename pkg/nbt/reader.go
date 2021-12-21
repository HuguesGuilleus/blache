// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2021, Hugues GUILLEUS
// All rights reserved.

package nbt

import (
	"io"
	"math"
)

// The reader
type reader []byte

func (r *reader) readByte() (b byte, err error) {
	if len(*r) == 0 {
		return 0, io.ErrUnexpectedEOF
	}
	b = (*r)[0]
	*r = (*r)[1:]
	return
}

func (r *reader) readInt16() (i int16, err error) {
	if len(*r) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	i = int16((*r)[0])<<8 +
		int16((*r)[1])

	*r = (*r)[2:]
	return
}

func (r *reader) readInt32() (i int32, err error) {
	if len(*r) < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	i = int32((*r)[0])<<24 +
		int32((*r)[1])<<16 +
		int32((*r)[2])<<8 +
		int32((*r)[3])

	*r = (*r)[4:]
	return
}

func (r *reader) readInt64() (i int64, err error) {
	if len(*r) < 8 {
		return 0, io.ErrUnexpectedEOF
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

func (r *reader) readFloat32() (float32, error) {
	i, err := r.readInt32()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(uint32(i)), nil
}

func (r *reader) readFloat64() (float64, error) {
	i, err := r.readInt64()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(uint64(i)), nil
}

func (r *reader) readString() (s string, err error) {
	var _lenStr int16
	_lenStr, err = r.readInt16()
	lenStr := int(_lenStr)
	if err != nil {
		return
	} else if len(*r) < lenStr {
		return "", io.ErrUnexpectedEOF
	}
	s = string((*r)[:lenStr])
	*r = (*r)[lenStr:]
	return
}
