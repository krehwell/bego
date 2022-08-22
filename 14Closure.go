package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func Closure() {
    fmt.Println("----- new int intSeq() initialize -----")

    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())


    fmt.Println("----- new int intSeq() initialize -----")

    newInt := intSeq()
    fmt.Println(newInt())
}
