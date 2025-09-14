// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package avltree

import (
	"encoding/json"

	"gitee.com/quant1x/gox/util/internal"
)

func assertSerializationImplementation() {
	var _ internal.JSONSerializer = (*Tree)(nil)
	var _ internal.JSONDeserializer = (*Tree)(nil)
}

// ToJSON outputs the JSON representation of the tree.
func (tree *Tree) ToJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	it := tree.Iterator()
	for it.Next() {
		elements[internal.ToString(it.Key())] = it.Value()
	}
	return json.Marshal(&elements)
}

// FromJSON populates the tree from the input JSON representation.
func (tree *Tree) FromJSON(data []byte) error {
	elements := make(map[string]interface{})
	err := json.Unmarshal(data, &elements)
	if err == nil {
		tree.Clear()
		for key, value := range elements {
			tree.Put(key, value)
		}
	}
	return err
}
