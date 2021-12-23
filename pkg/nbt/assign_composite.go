// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2021, Hugues GUILLEUS
// All rights reserved.

package nbt

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrorValueCanNotbeSet = errors.New("The value can not be set.")
	ErrorWrongType        = errors.New("The value can not be set with this type.")
)

func (r *reader) assignRoot(root reflect.Value) error {
	tagType, err := r.readByte()
	if err != nil {
		return fmt.Errorf("Read the tag type fail: %w", err)
	}

	if err := r.assignMapItem(root, tagType); err != nil {
		return fmt.Errorf("Assign to the root fail: %w", err)
	}
	return nil
}

// Assign a tagCompound to the value that can be a map[string]X or a struct.
func (r *reader) assignCompound(mapParent reflect.Value) error {
	mapType := mapParent.Type()
	if mapType.Kind() != reflect.Map {
		return fmt.Errorf("%w type:map[string]T, value type %s", ErrorWrongType, mapParent.String())
	} else if mapType.Key().Kind() != reflect.String {
		return fmt.Errorf("TagCompound can only be set to a map[string]T not map[%s]T", mapType.Key())
	}

	for {
		tagType, err := r.readByte()
		if err != nil {
			return fmt.Errorf("Read the tag type fail: %w", err)
		} else if tagType == tagEnd {
			return nil
		}

		if err := r.assignMapItem(mapParent, tagType); err != nil {
			return fmt.Errorf("Assign map[string]%s fail: %w", mapParent.Type().Elem(), err)
		}
	}
}

func (r *reader) assignMapItem(mapParent reflect.Value, tagType byte) error {
	tagName, err := r.readString()
	if err != nil {
		return fmt.Errorf("Read tag name fail: %w", err)
	}

	var itemValue reflect.Value
	switch elemType := mapParent.Type().Elem(); elemType.Kind() {
	case reflect.Map:
		itemValue = reflect.MakeMap(elemType)
	default:
		itemValue = reflect.Indirect(reflect.New(mapParent.Type().Elem()))
	}

	if err := r.assignValue(tagType, itemValue); err != nil {
		return fmt.Errorf("Assign value for the key %q fail: %w", tagName, err)
	}
	mapParent.SetMapIndex(reflect.ValueOf(tagName), itemValue)

	return nil
}

func (r *reader) assignValue(tagType byte, value reflect.Value) error {
	switch tagType {
	case tagByte:
		return r.assignByte(value)
	case tagInt16:
		return r.assignInt16(value)
	case tagInt32:
		return r.assignInt32(value)
	case tagInt64:
		return r.assignInt64(value)
	case tagFloat32:
		return r.assignFloat32(value)
	case tagFloat64:
		return r.assignFloat64(value)
	case tagString:
		return r.assignString(value)

	case tagCompound:
		return r.assignCompound(value)

	// Fail type
	case tagEnd:
		return fmt.Errorf("Unexpected Tag_End")
	default:
		return fmt.Errorf("Unknown this tag: %d", tagType)
	}
}
