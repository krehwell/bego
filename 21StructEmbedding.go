package main

import "fmt"

type base struct {
    num int
}

func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
    base
    str string
}

func StructEmbedding() {
    co := container{
        base: base{num: 1},
        str: "some name",
    }

    fmt.Println("co =", co)
    fmt.Println("co base num-", co.base.num)
    fmt.Println("co describe-", co.describe())

    type describer interface {
        describe() string
    }

    var d describer = &co
    fmt.Println("d describe-", d.describe())
}

