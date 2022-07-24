package main

import "fmt"

func Slices() {
    // slice is a `List` in generic pl
    s := make([]string, 3)
    fmt.Println("set", s)
    fmt.Println("length s", len(s))
    s[0] = "yakuza"
    s[1] = "yuza"
    s[2] = "suza"

    fmt.Println("set", s)
    fmt.Println("get", s[1])
    fmt.Println("len", len(s))

    fmt.Println("-----------")

    s = append(s, "marcopolo", "bolobolo")
    fmt.Println("after append(s, '...')", s)
    fmt.Println("len(s)", len(s))

    fmt.Println("-----------")

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("c is:", c)

    fmt.Println("-----------")

    l := c[2 : 3] // suza
    fmt.Println("l", l)

    l = c[: 5] // all c
    fmt.Println("l", l)

    fmt.Println("-----------")

    // initialize without using make
    t := []string{"bata", "subaru", "labu"}
    fmt.Println("t", t)

    fmt.Println("-----------")

    twoD := make([][]int, 3)
    fmt.Println("initial twoD", twoD)

    for i := 0; i < len(twoD); i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < len(twoD[i]); j++ {
            twoD[i][j] = i + j
        }
    }

    fmt.Println(twoD)
}
