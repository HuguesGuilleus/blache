// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2021, Hugues GUILLEUS
// All rights reserved.

package nbt

import (
	"fmt"
	"reflect"
)

/* Assign to simple value */

func (r *reader) assignByte(value reflect.Value) error {
	b, err := r.readByte()
	if err != nil {
		return fmt.Errorf("Read a byte fail: %w", err)
	}
	if !value.CanSet() {
		return fmt.Errorf("Can not set: %s", value.Type())
	}

	switch value.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value.SetInt(int64(b))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value.SetUint(uint64(b))
	default:
		return fmt.Errorf("Can not assign an byte to a %s", value.Type())
	}

	return nil
}

func (r *reader) assignString(value reflect.Value) error {
	s, err := r.readString()
	if err != nil {
		return err
	}

	if value.Type().Kind() != reflect.String {
		return fmt.Errorf("%w type:string to %s", ErrorWrongType, value.String())
	} else if value.CanSet() == false {
		return fmt.Errorf("%w %s", ErrorValueCanNotbeSet, value.String())
	}

	value.SetString(s)
	return nil
}
