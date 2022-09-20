package main

import (
    "fmt"
    "time"
)

func workerExample(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func ChannelSynchronization() {
    done := make(chan bool)

    go workerExample(done)

    <- done
}
