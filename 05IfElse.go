package main

import "fmt"

func IfElse() {
    if 7 % 2 == 0 {
        fmt.Println("it's even")
    } else {
        fmt.Println("it's odd")
    }

    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digit")
    }
}
