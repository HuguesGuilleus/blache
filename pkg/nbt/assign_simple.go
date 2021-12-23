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

	return assignInteger(value, int64(b))
}

func (r *reader) assignInt16(value reflect.Value) error {
	i, err := r.readInt16()
	if err != nil {
		return fmt.Errorf("Read int16 fail: %w", err)
	}
	return assignInteger(value, int64(i))
}

func (r *reader) assignInt32(value reflect.Value) error {
	i, err := r.readInt32()
	if err != nil {
		return fmt.Errorf("Read int32 fail: %w", err)
	}
	return assignInteger(value, int64(i))
}

func (r *reader) assignInt64(value reflect.Value) error {
	i, err := r.readInt64()
	if err != nil {
		return fmt.Errorf("Read int64 fail: %w", err)
	}
	return assignInteger(value, i)
}

func assignInteger(value reflect.Value, i int64) error {
	if !value.CanSet() {
		return fmt.Errorf("Can not set: %s", value.Type())
	}

	switch value.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value.SetUint(uint64(i))
	default:
		return fmt.Errorf("Can not assign an byte to a %s", value.Type())
	}
	return nil
}

func (r *reader) assignFloat32(value reflect.Value) error {
	f, err := r.readFloat32()
	if err != nil {
		return fmt.Errorf("Read float32 fail: %w", err)
	}
	return assignFloat(value, float64(f))
}

func (r *reader) assignFloat64(value reflect.Value) error {
	f, err := r.readFloat64()
	if err != nil {
		return fmt.Errorf("Read float64 fail: %w", err)
	}
	return assignFloat(value, f)
}

func assignFloat(value reflect.Value, f float64) error {
	if !value.CanSet() {
		return fmt.Errorf("Can not set: %s", value.Type())
	}

	switch value.Type().Kind() {
	case reflect.Float32, reflect.Float64:
		value.SetFloat(f)
		return nil
	default:
		return fmt.Errorf("Can not assign an byte to a %s", value.Type())
	}
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
