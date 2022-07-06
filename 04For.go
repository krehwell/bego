package main

import (
    "fmt"
)

func ForLoop() {
    i := 0

    for i <= 3 {
        fmt.Println(i)
        i += 1
    }

    fmt.Println("----------")

    for i = 0; i < 10; i++ {
        fmt.Println(i)
    }

    fmt.Println("----------")

    for {
        fmt.Println("This is just while loop")
        break
    }
}
