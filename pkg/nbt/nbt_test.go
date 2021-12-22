// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2021, Hugues GUILLEUS
// All rights reserved.

package nbt

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed hello_world.nbt
var helloNBT []byte

func TestHello(t *testing.T) {
	var mapValue = make(map[string]map[string]string)
	assert.NoError(t, Unmarshal(helloNBT, mapValue))
	assert.Equal(t, map[string]map[string]string{
		"hello world": map[string]string{
			"name": "Bananrama",
		},
	}, mapValue)
}
