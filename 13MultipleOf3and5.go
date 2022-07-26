package main

import "fmt"

func MultipleOf3And5() {
    /**
     * If we list all the natural numbers below 10 that are multiples of 3 or 5,
     * we get 3, 5, 6 and 9. The sum of these multiples is 23.
     * Find the sum of all the multiples of 3 or 5 below 1000
     * answer: 233168
     */
    const MAX_N = 1_000
    bucket := []int{}

    for i := 1; i < MAX_N; i++ {
        if i % 5 == 0 || i % 3 == 0 {
            bucket = append(bucket, i)
        }
    }

    sum := 0
    for i := 0; i < len(bucket); i++ {
        sum += bucket[i]
    }

    fmt.Println(sum)
}
