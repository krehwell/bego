package main

import "fmt"


func Maps() {
    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 0

    fmt.Println(m)

    v1 := m["k1"]
    fmt.Println("v1 is = ", v1)

    fmt.Println("length of m = ", len(m))

    delete(m, "k1")

    fmt.Println("length of m after delete k1 = ", len(m), "-- and m = ", m)

    result, prs := m["k1"]
    fmt.Println("result, prs := m['k1'] |", result, prs)

    // inline initialization
    n := map[string]string{"child1": "yuza", "child2":"yakuza"}
    fmt.Println("inline initialization = ", n)
}
