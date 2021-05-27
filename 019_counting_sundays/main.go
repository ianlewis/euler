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
)

type month int

const (
	January = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

type weekday int

const (
	Monday weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func daysInMonth(m month, y int) int {
	// Special case for February to account for leap years.
	if m == February {
		if y%4 == 0 && !(y%100 == 0 && y%400 != 0) {
			return 29
		}
		return 28
	}
	d := []int{
		31, // Jan
		0,  // Feb (not used)
		31, // Mar
		30, // Apr
		31, // May
		30, // Jun
		31, // Jul
		31, // Aug
		30, // Sep
		31, // Oct
		30, // Nov
		31, // Dec
	}
	return d[int(m)-1]
}

func startOfNextMonth(startWeekday weekday, m month, y int) weekday {
	return weekday((int(startWeekday) + daysInMonth(m, y)) % 7)
}

func main() {
	// I calculate the number of Sundays at the start of the month by
	// calculating the starting weekday for each subsequent month. I don't use
	// the time package because I just thought that would be cheating ;)

	// Find the start weekday of Jan 1, 1901 given that Jan 1, 1900 starts on
	// Monday.
	startWeekday := Monday
	for m := 1; m <= 12; m++ {
		startWeekday = startOfNextMonth(startWeekday, month(m), 1900)
	}

	// Calculate the number of Sundays on the first of the month.
	noSundays := 0
	for y := 1901; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			if startWeekday == Sunday {
				noSundays++
			}
			startWeekday = startOfNextMonth(startWeekday, month(m), y)
		}
	}
	fmt.Println(noSundays)
}
