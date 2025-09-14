// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	dll "gitee.com/quant1x/gox/util/doublylinkedlist"
	"gitee.com/quant1x/gox/util/internal"
)

// DoublyLinkedListExample to demonstrate basic usage of DoublyLinkedList
func main() {
	list := dll.New()
	list.Add("a")                         // ["a"]
	list.Append("b")                      // ["a","b"] (same as Add())
	list.Prepend("c")                     // ["c","a","b"]
	list.Sort(internal.StringComparator)  // ["a","b","c"]
	_, _ = list.Get(0)                    // "a",true
	_, _ = list.Get(100)                  // nil,false
	_ = list.Contains("a", "b", "c")      // true
	_ = list.Contains("a", "b", "c", "d") // false
	list.Remove(2)                        // ["a","b"]
	list.Remove(1)                        // ["a"]
	list.Remove(0)                        // []
	list.Remove(0)                        // [] (ignored)
	_ = list.Empty()                      // true
	_ = list.Size()                       // 0
	list.Add("a")                         // ["a"]
	list.Clear()                          // []
}
