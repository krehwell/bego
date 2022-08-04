package main

import "fmt"

func trialDivision(n int) []int {
    a := []int{}
    f := 2

    for ;n > 1; {
        if n % f == 0 {
            a = append(a, f)
            n /= f
        } else {
            f += 1
        }
    }

    return a
}

func LargestPrimeNumber() {
    /**
     * The prime factors of 13195 are 5, 7, 13 and 29.
     * What is the largest prime factor of the number 600851475143?”
     * answer: 6857
     */
     result := trialDivision(600851475143)
     fmt.Println(result, result[len(result) - 1])
}
