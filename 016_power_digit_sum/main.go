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
	"math/big"
	"strconv"
)

// powerDigitSum calculates 2^e as a big integer and then sum's the digits of
// that number.
func powerDigitSum(e int) int {
	n := big.NewInt(2)
	for i := 0; i < e-1; i++ {
		n = n.Mul(n, big.NewInt(2))
	}

	sum := 0
	for _, c := range n.String() {
		d, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		sum += d
	}

	return sum
}

func main() {
	fmt.Println(powerDigitSum(1000))
}
