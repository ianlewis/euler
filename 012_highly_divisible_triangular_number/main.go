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

// The number of divisiors to look for.
const num = 500

// numDivisors returns the number of divisors for the given number. It is
// assumed a divisor is a natural number that divides the given number cleanly.
func numDivisors(n int) int {
	num := 0

	// We can check the numbers up to sqrt and add both products.
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			if n/i == i {
				// If this is a sqrt only add it as a single divisor.
				num++
			} else {
				// Otherwise add two numbers as both products are divisors.
				num += 2
			}
		}
	}
	return num
}

func main() {
	i := 1

	// t is the current triangular number. 1 is the 1st triangular number.
	t := 1
	for {
		if numDivisors(t) > num {
			fmt.Println(t)
			break
		}

		i++
		// We keep track of the current triangular number by just adding the i to
		// the previous triangular number to save on computation.
		t += i
	}
}
