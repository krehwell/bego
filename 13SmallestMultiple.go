package main

import "fmt"

func SmallestMultiple() {
    /**
     * 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
     * What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
     * answer: 232792560
     */
     min := 1
     max := 20

     numToDivide := max

     for i := min; i <= max; {
         if numToDivide % i != 0 {
             numToDivide += 1
             i = min
             continue
         } else {
             i += 1
         }

         if i >= max {
             break
         }
     }

    fmt.Println(numToDivide)
}
