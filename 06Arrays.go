package main

import ("fmt")

func Arrays() {
    var a [5]int

    fmt.Println("a:", a)

    a[4] = 100

    fmt.Println("a should be changed now", a)
    fmt.Println("length of a", len(a))

    fmt.Println("----------")

    b := [5]int{1, 2, 3, 4}
    fmt.Println("b:", b)

    fmt.Println("----------")

    // 2D array
    var c [4][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 2; j++ {
            c[i][j] = i + j
        }
    }

    fmt.Println("c:", c)
    fmt.Println("length of c", len(c))
    fmt.Println("length of c0", len(c[0]))
}
