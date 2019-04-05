// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedliststack

import (
	"github.com/mymmsc/gox/util"
	"github.com/mymmsc/gox/util/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*util.Stack)(nil)
	var _ containers.JSONDeserializer = (*util.Stack)(nil)
}

// ToJSON outputs the JSON representation of the stack.
func (stack *util.Stack) ToJSON() ([]byte, error) {
	return stack.list.ToJSON()
}

// FromJSON populates the stack from the input JSON representation.
func (stack *util.Stack) FromJSON(data []byte) error {
	return stack.list.FromJSON(data)
}
