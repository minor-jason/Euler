// euler1.go
// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23.
// 
// Find the sum of  all the multiples of 3 or 5 below 1000.
package main

import (
	"fmt"
	"math"
	)
	
func main() {
	var x, sum float64
	sum = 0	
	x = 0

	for i := 1; i < 1000; i++ {
		x = x + 1
		if math.Mod(x, 3) == 0 || math.Mod(x, 5) == 0 {
			sum = sum + x
		}
	}
	fmt.Printf("The sum is %v\n", sum )
}
