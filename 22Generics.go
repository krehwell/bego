package main

import "fmt"

type element[T any] struct {
    next *element[T]
    val T
}

type List[T any] struct {
    head, currPointer *element[T]
}

func (lst *List[T]) Push(v T) {
    if lst.head == nil {
        lst.head = &element[T]{val: v}
        lst.currPointer = lst.head
    } else {
        lst.currPointer.next = &element[T]{val: v}
        lst.currPointer = lst.currPointer.next
    }
}

func (lst *List[T]) GetAll() []T {
    vals := []T{}

    for e := lst.head; e != nil; e = e.next {
        vals = append(vals, e.val)
    }

    return vals
}

func Generics() {
    list := List[int]{}

    list.Push(10)
    list.Push(12)
    list.Push(16)
    list.Push(14)

    fmt.Println(list.GetAll())
}
