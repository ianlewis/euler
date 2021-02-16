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

const longNumStr = `7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450`

const num = 13

func main() {
	// Create a rotating array of adjacent numbers.
	digits := make([]int, num)
	// product is the greatest product found so far.
	product := 0
	for i := range longNumStr {
		// Read each digit
		d, err := strconv.Atoi(string(longNumStr[i]))
		if err != nil {
			panic(err)
		}
		// Store the digit at the end of the array.
		digits[num-1] = d

		// Calculate the product of the adjacent digits while rotating the
		// digits in the digits array.
		// For i < num we know that values in digits is initialized to zero
		// and will always result in a lower product than those seen before.
		newProduct := 1
		for j := range digits {
			newProduct *= digits[j]
			if j+1 < len(digits) {
				digits[j] = digits[j+1]
			}
		}
		if newProduct > product {
			product = newProduct
		}
	}
	fmt.Println(product)
}
