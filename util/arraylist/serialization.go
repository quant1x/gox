// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arraylist

import (
	"encoding/json"
	"gitee.com/quant1x/gox/util/internal"
)

func assertSerializationImplementation() {
	var _ internal.JSONSerializer = (*List)(nil)
	var _ internal.JSONDeserializer = (*List)(nil)
}

// ToJSON outputs the JSON representation of list's elements.
func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list.elements[:list.size])
}

// FromJSON populates list's elements from the input JSON representation.
func (list *List) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &list.elements)
	if err == nil {
		list.size = len(list.elements)
	}
	return err
}
