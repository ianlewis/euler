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

const max = 2000000

func isPrime(value int64) bool {
	for i := int64(2); i <= int64(math.Sqrt(float64(value+1))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Sum of the primes found so far.
	sum := int64(2)

	// Iterate by two because we know that all primes starting with 3 must
	// be odd.
	for i := int64(3); i < max; i += 2 {
		if isPrime(i) {
			sum += i
		}

		i += 2
	}
	fmt.Println(sum)
}
