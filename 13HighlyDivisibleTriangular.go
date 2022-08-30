package main

import (
	"fmt"
	"math"
	// "math"
)

func HighlyDivisibleTriangular() {
    /**
     * The sequence of triangle numbers is generated by adding the natural numbers.
     * So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28. The first ten terms would be:
     * 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, …
     * Let us list the factors of the first seven triangle numbers:
     * 1: 1
     * 3: 1,3
     * 6: 1,2,3,6
     * 10: 1,2,5,10
     * 15: 1,3,5,15
     * 21: 1,3,7,21
     * 28: 1,2,4,7,14,28
     * We can see that 28 is the first triangle number to have over five divisors.
     * What is the value of the first triangle number to have over five hundred divisors?
     * for first to have 300 divisor it is on 2162160
     * answer: 76576500
     */
    triangle := 0
    for n := 1; ; n++ {
        divisorsCount := 0

        triangle += n
        to := int(math.Sqrt(float64(triangle)))

        for i := 1; i <= to; i++ {
            if triangle % i == 0 {
                divisorsCount += 2
            }
        }


        if divisorsCount >= 500 {
            fmt.Println("found at", triangle)
            break
        }
    }
}
