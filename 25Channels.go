package main

import (
    "fmt"
    "time"
)

func Channels() {
    message := make(chan string)

    go func() {
        time.Sleep(time.Second * 2)
        message <- "ping"
    }()

    fmt.Println("Boo")
    fmt.Println(<-message)
}
