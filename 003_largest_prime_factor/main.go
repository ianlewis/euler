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

func main() {
	n := int64(600851475143)
	largest := int64(0)
	// Divide out all 2s
	for n%2 == 0 {
		n /= 2
		fmt.Println(2)
		largest = 2
	}

	// Divide out all other primes. Prime factors must be odd numbers less than
	// the sqrt of N.
	sqrtN := int64(math.Sqrt(float64(n)))
	for i := int64(3); i < sqrtN; i += 2 {
		for n%i == 0 {
			n /= i
			fmt.Println(i)
			largest = i
		}
	}

	// Catch the case that the remaining factor is greater than cases so far.
	if n > largest {
		largest = n
	}
	fmt.Println(largest)
}
