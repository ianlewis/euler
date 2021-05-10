// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
)

func main() {
	var longestStart int64
	var longestChain int64
	for n := int64(1); n < int64(1000000); n++ {
		i := n
		size := int64(1)
		for i != 1 {
			if i%2 == 0 {
				i = i / 2
			} else {
				i = 3*i + 1
			}
			size++
		}

		if size > longestChain {
			longestStart = n
			longestChain = size
		}
	}
	fmt.Printf("%d produces a chain of size %d\n", longestStart, longestChain)
}
