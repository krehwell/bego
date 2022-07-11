/**
 * Make a 2D slices.
 * make it so that it will be [[0], [1, 2], [3, 4, 5], [6, 7, 8, 9]]
 */

package main

import "fmt"

func SlicesTest() {
    s := make([][]int, 4)

    fmt.Println(s)

    for i, l := 0, 0; i < 4; i++ {
        s[i] = make([]int, i + 1)

        for j := 0; j < len(s[i]); j++ {
            s[i][j] = l
            l++
        }
    }

    fmt.Println(s)
}


