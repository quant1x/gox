// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binaryheap

import (
	"github.com/mymmsc/gox/util"
	"github.com/mymmsc/gox/util/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*util.Heap)(nil)
	var _ containers.JSONDeserializer = (*util.Heap)(nil)
}

// ToJSON outputs the JSON representation of the heap.
func (heap *util.Heap) ToJSON() ([]byte, error) {
	return heap.list.ToJSON()
}

// FromJSON populates the heap from the input JSON representation.
func (heap *util.Heap) FromJSON(data []byte) error {
	return heap.list.FromJSON(data)
}
