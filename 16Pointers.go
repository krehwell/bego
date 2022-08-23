package main

import "fmt"

func zeroval(val int) {
    val = 0
}

func zeroptr(val *int) {
    *val = 0
}

func Pointers() {
    i := 2
    fmt.Println("initial i :=", i)

    fmt.Println("----- after call zeroval")
    zeroval(i)
    fmt.Println(i)

    fmt.Println("----- after call zeroptr")
    zeroptr(&i)
    fmt.Println(i)

    fmt.Println("address of i =", &i)
}
