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

import "fmt"

func main() {
	n := int64(600851475143)
	for i := n / 2; i > 0; i-- {
		if n%i == 0 {
			// i is a factor. Find if i is prime
			prime := true
			for j := int64(0); j < i; j++ {
				if i%j == 0 {
					// not prime
					prime = false
					break
				}
			}
			if prime {
				fmt.Println(i)
				return
			}
		}
	}
	// Largest prime factor of n is n itself
	fmt.Println(n)
}
