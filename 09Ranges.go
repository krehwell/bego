package main

import "fmt"

func Ranges() {
    nums := []int{1,2,3,4,5}

    sum := 0

    for _, num := range nums {
        sum += num
    }

    fmt.Println("result from range sum =", sum)

    for i, num := range nums {
        sum += i + num
    }

    fmt.Println("result from range sum + i in each iteration =", sum)

    fmt.Println("-------")

    kvs := map[string]string{"a": "apple", "b": "banana"}

    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key of kvs = ", k)
    }

    fmt.Println("------- range over a string -------")

    for i, k := range "go" {
        fmt.Println(i, k)
    }
}
