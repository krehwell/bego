package main

import (
    "fmt"
    "time"
)

func ChannelSelect() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        for ; ; {
            time.Sleep(time.Millisecond * 500)
            c1<- "one"
        }
    }()

    go func() {
        for ; ; {
            time.Sleep(time.Second * 2)
            c2<- "two"
        }
    }()

    for ; ; {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
