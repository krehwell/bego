package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(2 * time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func WaitGroups() {
    var wg sync.WaitGroup

    for i := 0; i <= 5; i++ {
        wg.Add(1)

        i := i

        go func() {
            defer wg.Done()
            worker(i)
        }()
    }

    wg.Wait()
}
