// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedliststack

import "github.com/mymmsc/gox/util"

func assertSerializationImplementation() {
	var _ util.JSONSerializer = (*Stack)(nil)
	var _ util.JSONDeserializer = (*Stack)(nil)
}

// ToJSON outputs the JSON representation of the stack.
func (stack *Stack) ToJSON() ([]byte, error) {
	return stack.list.ToJSON()
}

// FromJSON populates the stack from the input JSON representation.
func (stack *Stack) FromJSON(data []byte) error {
	return stack.list.FromJSON(data)
}
