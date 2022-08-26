package main

import "fmt"

type Rect struct {
    width, height int
}

func (r *Rect) area() int {
    return r.width + r.height
}

func (r Rect) perim() int {
    return r.width*2 + r.height*2
}

func MethodExample() {
    r := Rect{width: 20, height: 32}
    fmt.Println("r - area()", r.area())

    rb := &r
    fmt.Println("rb - perim()", rb.perim())
}
