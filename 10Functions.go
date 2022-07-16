package main

import "fmt"

func plus(a int, b int) int {
    return a + b
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func Functions() {
    result0 := plus(1, 4)
    result1 := plusPlus(1, 4, 5)
    fmt.Println(result0, result1)
}
