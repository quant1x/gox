// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/quant1x/gox/util/internal"
)

// SortExample to demonstrate basic usage of basic sort
func main() {
	strings := []interface{}{}                        // []
	strings = append(strings, "d")                    // ["d"]
	strings = append(strings, "a")                    // ["d","a"]
	strings = append(strings, "b")                    // ["d","a",b"
	strings = append(strings, "c")                    // ["d","a",b","c"]
	internal.Sort(strings, internal.StringComparator) // ["a","b","c","d"]
}
