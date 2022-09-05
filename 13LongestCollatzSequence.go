package main

import "fmt"

func onOdd(n int) int {
    return (3 * n) + 1
}

func onEven(n int) int {
    return n / 2
}

func LongestCollatzSequence() {
    MAX_TO_BRUTE := 1_000_000

    maxChain := 0

    at := 2;

    for i := 2; i <= MAX_TO_BRUTE; i++ {
        count := 0
        for currNum := i; currNum > 1; {
            if currNum % 2 == 0 {
                currNum = onEven(currNum)
            } else {
                currNum = onOdd(currNum)
            }

            count++
        }

        if count > maxChain {
            at = i
            maxChain = count
        }
    }

    fmt.Println(at, maxChain + 1) // +1 because including `i`
}
