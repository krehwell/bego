package main

import (
	"fmt"
	"math/big"
)

func getFactorial(n int) *big.Int {
    sum := big.NewInt(1)
    for i := n; i > 0; i-- {
        sum = sum.Mul(sum, big.NewInt(int64(i)))
    }

    return sum
}

func LatticePath() {
    /**
     * Starting in the top left corner of a 2×2 grid,
     * and only being able to move to the right and down, there are exactly 6 routes
     * to the bottom right corner.
     * How many such routes are there through a 20×20 grid?
     * formula is: result = (2n)! / n!^2
     * answer: 137846528820
     */
    n := 20

    topFormula := getFactorial(2 * n)
    _belowFormula := getFactorial(n)
    belowFormula := big.NewInt(0).Mul(_belowFormula, _belowFormula)

    result := big.NewInt(0).Div(topFormula, belowFormula)

    fmt.Println("result", result)
}
