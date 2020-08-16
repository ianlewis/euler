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
	"strconv"
)

func isPalindrome(x int) bool {
	y := string(strconv.Itoa(x))
	z := ""
	for i := len(y) - 1; i >= 0; i-- {
		z += string(y[i])
	}
	return z == y
}

func main() {
	a := 0
	b := 0
	largest := 0
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			x := i * j
			if isPalindrome(x) && x > largest {
				a = i
				b = j
				largest = x
			}
		}
	}
	fmt.Println(a, b, largest)
}
