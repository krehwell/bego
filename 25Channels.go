package main

import (
    "fmt"
    "time"
)

func AddOne(ch chan int, val int) {
    val += 1
    time.Sleep(time.Second * 2)
    ch <- val
}

func MulTen(ch chan int, resChan chan int) {
    val := <-ch
    val *= 10
    resChan <- val
}

func Channels() {
    ch := make(chan int)
    resChan := make(chan int)

    go AddOne(ch, 9)
    go MulTen(ch, resChan)

    fmt.Println(<-resChan)
}
