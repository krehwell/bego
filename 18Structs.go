package main

import "fmt"

type Person struct {
    name string
    age int
}

func newPerson(name string) *Person {
    p := Person{name: name}
    p.age = 42
    return &p
}

func StructInGo() {
    fmt.Println(Person{"alice", 42}) // no need specify prop
    fmt.Println(Person{name: "yuza", age: 22})
    fmt.Println(Person{name: "yakuza"}) // age will be 0
    fmt.Println(&Person{"yakuza", 69})
    fmt.Println(newPerson("new yakuza constructor"))

    fmt.Println("----- start initializing -----")
    s := Person{"yuza", 99999}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.name)

    sp.age = 42
    fmt.Println(s.age)
}
