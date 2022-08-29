package main

// lint:ignore U1000 Ignore unused function temporarily for debugging
import "fmt"

func MapKeys[K comparable, V any](items map[K]V) []K {
    keys := []K{}

    for key := range items {
        keys = append(keys, key)
    }

    return keys
}

type element[T any] struct {
    next *element[T]
    val T
}

type List[T any] struct {
    head, tail *element[T]
}

func (list *List[T]) Push(val T) {
    if list.tail == nil {
        list.head = &element[T]{val: val}
        list.tail = list.head
    } else {
        list.tail.next = &element[T]{val: val}
        list.tail = list.tail.next
    }
}

func (list *List[T]) GetAllItems() []T {
    vals := []T{}

    for e := list.head; e != nil; e = e.next {
        vals = append(vals, e.val)
    }

    return vals
}

func Generics() {
    fmt.Println("----- MapKeys -----")
    object := map[int]string{10: "one", 40: "two", 69: "three"}
    fmt.Println(MapKeys(object))

    fmt.Println("----- Linked List -----")
    l := List[int]{}
    l.Push(10)
    l.Push(12)
    l.Push(14)
    l.Push(18)
    fmt.Println(l.GetAllItems())
}
