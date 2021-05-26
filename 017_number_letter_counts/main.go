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

// digits are the numbers 0-9 by index
var digits = [10]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

// teens are values 10-19 indexed by x-10
var teens = [10]string{
	"ten",       // 0+10
	"eleven",    // 1+10
	"twelve",    // 2+10
	"thirteen",  // 3+10
	"fourteen",  // 5+10
	"fifteen",   // 5+10
	"sixteen",   // 6+10
	"seventeen", // 7+10
	"eighteen",  // 8+10
	"ninteen",   // 9+10
}

// tens are multiples of 10 indexed by x/10
var tens = [10]string{
	"",        // undefined
	"ten",     // 1*10
	"twenty",  // 2*10
	"thirty",  // 3*10
	"fourty",  // 4*10
	"fifty",   // 5*10
	"sixty",   // 6*10
	"seventy", // 7*10
	"eighty",  // 8*10
	"ninety",  // 9*10
}

// powers are the powers of 10
var powers = [7]string{
	"zero",     // 10^0
	"ten",      // 10^1
	"hundred",  // 10^2
	"thousand", // 10^3
	"",         // 10^4
	"",         // 10^5
	"million",  // 10^6
}

// numToString converts a number to a string in plain English.
func numToString(n int) string {
	str := ""

	// Remove powers of ten
	for i := len(powers) - 1; i > 1; i-- {
		// Powers like 10^4,10^5 are special because they convert to "X
		// thousand" and the like.
		if powers[i] == "" {
			continue
		}
		p := int(math.Pow(float64(10), float64(i)))
		x := n / p
		if x > 0 {
			if str != "" {
				str += " "
			}
			str += fmt.Sprintf("%s %s", numToString(x), powers[i])
			n -= x * p
		}
	}

	// If we need to add an "and"
	if str != "" && n > 0 {
		str += " and"
	}

	// Remove multiples of ten > 20
	for i := len(tens); i > 1; i-- {
		p := 10 * i
		if n >= p {
			if str != "" {
				str += " "
			}
			str += tens[i]
			n -= p
			break
		}
	}

	// Convert numbers 0-20
	if n >= 10 {
		for i := len(teens) - 1; i >= 0; i-- {
			if n == 10+i {
				if str != "" {
					str += " "
				}
				str += teens[i]
				n -= 10 + i
				break
			}
		}
	} else {
		if n > 0 {
			if str != "" {
				str += " "
			}
			str += digits[n]
		} else {
			if str == "" {
				str = digits[n]
			}
		}
	}

	return str
}

// numLetters calculates the number of letters in the string.
func numLetters(str string) int {
	letters := 0
	for _, r := range str {
		if r >= 'a' && r <= 'z' {
			letters++
		}
	}
	return letters
}

func main() {
	letters := 0
	for i := 1; i <= 1000; i++ {
		// Convert the number to string and then calculate the number of
		// letters. It's done this way mostly for debuggability.
		letters += numLetters(numToString(i))
	}
	fmt.Println(letters)
}
