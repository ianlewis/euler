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

// There are a couple ways to solve this. The simple solution is to just find
// all the paths for a 20x20 grid but the number grows very big very quickly. A
// faster and more general solution would be to use dynamic programming to
// build up the number of paths available for every size grid up to 20x20.
//
// Each vertex in the grid can hold the number of paths to the end from that
// vertex. This number is equal to the number from the vertex below it plus the
// vertex to the right of it. We can work our way up the grid calculating the
// number of paths at each point. However, since the grid is symetrical along the
// diagonal, we only need to maintain the current row we are processing and the
// previous row.
func latticeDynamic(w, h int) int {
	grid := make([][2]int, w+1)

	// initialize the array. We store locations in the reverse order from the
	// images in the problem so our paths will go bottom right to top left. The
	// final column position in the path is at grid[0]
	// 0, 1, 1, 1, ...
	grid[0][1] = 0
	for i := 1; i < w+1; i++ {
		grid[i][1] = 1
	}

	// Rather than swap the two slices we just swap the index by using j%2.
	j := 0
	for ; j < h+1; j++ {
		grid[0][j%2] = 1 // the leading edge
		for i := 1; i < w+1; i++ {
			grid[i][j%2] = grid[i-1][j%2] + grid[i][(j+1)%2]
		}
	}

	return grid[w][j%2]
}

func main() {
	fmt.Println(latticeDynamic(20, 20))
}
