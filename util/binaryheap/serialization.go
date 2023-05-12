// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binaryheap

import (
	"gitee.com/quant1x/gox/util"
)

func assertSerializationImplementation() {
	var _ util.JSONSerializer = (*Heap)(nil)
	var _ util.JSONDeserializer = (*Heap)(nil)
}

// ToJSON outputs the JSON representation of the heap.
func (heap *Heap) ToJSON() ([]byte, error) {
	return heap.list.ToJSON()
}

// FromJSON populates the heap from the input JSON representation.
func (heap *Heap) FromJSON(data []byte) error {
	return heap.list.FromJSON(data)
}
