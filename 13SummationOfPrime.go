package main

import (
    "fmt"
    "math"
)

func isPrime_(n int) bool {
    if n < 2 {
        return false
    }

    max := math.Sqrt(float64(n))
    for i := 2; i <= int(max); i++ {
        if n % i == 0 {
            return false
        }
    }

    return true
}

func SummationOfPrime() {
    /**
     * The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
     * Find the sum of all the primes below two million.
     * answer: 142913828922
     */
    sum := 0

    for i := 2; i < 2_000_000; i++ {
        if isPrime_(i) {
            sum += i
        }
    }

    fmt.Println(sum)
}
