package main

import "fmt"

func SumSquareDifference() {
    /**
     * Hence the difference between the sum of the squares of the first
       ten natural numbers and the square of the sum is .
     * Find the difference between the sum of the squares of the
       first one hundred natural numbers and the square of the sum
     * answer: 25164150
     */
    min := 1
    max := 100

    resultOfMinToMaxSquares := 0
    resultOfSumMinMaxSquares := 0

    for i := min; i <= max; i++ {
        resultOfMinToMaxSquares += i * i
        resultOfSumMinMaxSquares += i
    }

    resultOfSumMinMaxSquares = resultOfSumMinMaxSquares * resultOfSumMinMaxSquares

    fmt.Println("resultOfMinToMaxSquares", resultOfMinToMaxSquares)
    fmt.Println("resultOfSumMinMaxSquares", resultOfSumMinMaxSquares)
    fmt.Println("answer: ", resultOfSumMinMaxSquares - resultOfMinToMaxSquares)
}
