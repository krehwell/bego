package main

import (
    "fmt"
    "time"
)

func workerJob(id int, jobs <-chan int, results chan<- int) {
    fmt.Println("num of jobs injected", len(jobs))

    for j := range jobs {
        fmt.Println("worker", id, "started job", j)
        time.Sleep(time.Second * 1)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func WorkerPools() {
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    for w := 0; w < numJobs; w++ {
        go workerJob(w, jobs, results)
    }

    for j := 0; j < numJobs; j++ {
        jobs <- j
    }

    for a := 0; a < numJobs; a++ {
        <-results
    }
}


