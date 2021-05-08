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
	"math"
)

const sum = 1000

func main() {
	// It's unclear whether natural numbers should start at zero or one but
	// based on the description that a < b < c, the only number that could be zero
	// is a.
	// Loop from 2 to 999 because the lowest triplet could only be 0+1+2 and
	// the highest c could be would be in the triplet 0+1+999.
	for c := 2; c < sum-1; c++ {
		// Loop starting at one because a < b and if a = 0 then b must be 1.
		for b := 1; b < c; b++ {
			a := math.Sqrt(float64(c*c - b*b))
			// Verify that:
			// 1. a is a natural number
			// 2. a < b (b is already < c)
			// c. a+b+c == 1000
			if a == math.Trunc(a) && int(a) < b && int(a)+b+c == sum {
				fmt.Printf("%d + %d + %d = %d\n", int(a), b, c, sum)
				break
			}
		}
	}
}
