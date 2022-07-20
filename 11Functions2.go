package main

import "fmt"

func Get2Values() (int, int) {
    return 4, 5
}

func Functions2() {
    a, b := Get2Values()

    fmt.Println(a, b)

    _, c := Get2Values()
    fmt.Println(c)
}
