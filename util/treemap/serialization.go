// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treemap

import (
	"github.com/mymmsc/gox/util"
	"github.com/mymmsc/gox/util/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*util.Map)(nil)
	var _ containers.JSONDeserializer = (*util.Map)(nil)
}

// ToJSON outputs the JSON representation of the map.
func (m *util.Map) ToJSON() ([]byte, error) {
	return m.tree.ToJSON()
}

// FromJSON populates the map from the input JSON representation.
func (m *util.Map) FromJSON(data []byte) error {
	return m.tree.FromJSON(data)
}
