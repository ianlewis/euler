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

// sumDivisors returns the sum of divisors for the given number. It is
// assumed a divisor is a natural number that divides the given number cleanly.
func sumDivisors(n int) int {
	// Include one as a divisor but not the number itself.
	sum := 1
	// We can check the numbers up to sqrt and add both products.
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			if n/i == i {
				// If this is a sqrt only add it as a single divisor.
				sum += i
			} else {
				sum += i
				sum += n / i
			}
		}
	}
	return sum
}

func main() {
	sum := 0
	for i := 1; i < 10000; i++ {
		s := sumDivisors(i)
		if i != s && sumDivisors(s) == i {
			sum += i
		}
	}
	fmt.Println(sum)
}
