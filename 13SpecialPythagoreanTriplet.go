package main

import (
    "fmt"
)

func SpecialPythagoreanTriplet() {
    /**
     * Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
     * For example, 3^2 + 4^2 = 5^2 || 9 + 16 = 25
     * There exists exactly one Pythagorean triplet for which a + b + c = 1000
     * answer: 200 + 375 + 425 = 1000
     */
     target := 1000
     result := [3]int{}

     out:
     for a := 1; a <= target; a++ {
         for b := 1; b <= target; b++ {
             for c := 1; c <= target; c++ {
                 isPythagoras := a*a + b*b == c*c
                 if isPythagoras {
                     sum := a + b + c
                     if sum == target {
                         result = [3]int{a, b, c}
                         break out
                     }
                 }
             }
         }
     }

     fmt.Println(result)
}
