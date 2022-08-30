package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func Goroutine() {
    f("direct")

    go f("thread")

    go func(msg string) {
        fmt.Println(msg)
    } ("anonymous call")

    time.Sleep(time.Second)
    fmt.Println("done")
}
