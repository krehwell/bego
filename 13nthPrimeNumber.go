package main

import (
    "fmt"
    "math"
)

func isPrime(n int) bool {
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

func NthPrimeNumber() {
    /**
     *
     * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
     * What is the 10 001st prime number?
     * answer: 104743
     */
    n := 10_001
    primeCounter := 0

    for i := 2; primeCounter != n; i++ {
        if isPrime(i) {
            primeCounter++
            fmt.Println("result is:", i)
        }
    }

    fmt.Println(primeCounter)
}
