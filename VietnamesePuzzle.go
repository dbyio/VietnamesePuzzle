//
// Solving The Gardian's math puzzle for vietnamese 3rd graders (efficiently)
// Golang implementation
//
// _ + 13 * _ / _ + _ + 12 * _ - _ - 11 + _ * _ / _ - 10 = 66
//
// Author: nperraud <np@bitbox.io>
//

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var solutions [][]int
	scope := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // all solutions are permutations of this

	start := time.Now()
	solutions = unpuzzle(scope)
	end := time.Now()
	for _, i := range solutions {
		fmt.Println(i)
	}
	fmt.Printf("\n-- %d solutions found in %f seconds. --\n", len(solutions), end.Sub(start).Seconds())
	return
}

func snake_it(x []int) bool {
	if Round((float64)(x[0])+(13*(float64)(x[1])/(float64)(x[2]))+
		(float64)(x[3])+(12*(float64)(x[4]))-(float64)(x[5])-11+
		((float64)(x[6])*(float64)(x[7])/
			(float64)(x[8]))-10, 0.000001, 2) == 66 {
		return true
	}
	return false
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

// Unpuzzle the snake
// Implement Heap's algorithm for permutation and test as we go
func unpuzzle(scope []int) [][]int {
	res := [][]int{}
	var generate func(int, []int)

	generate = func(n int, arr []int) { // Heap's algorithm
		if n == 1 {
			if snake_it(arr) == true {
				A := make([]int, len(arr))
				copy(A, arr)
				res = append(res, A)
			}
		} else {
			for i := 0; i < n; i++ {
				generate(n-1, arr)
				if n%2 == 0 {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				} else {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				}
			}
		}
	}
	generate(len(scope), scope)
	return res
}
