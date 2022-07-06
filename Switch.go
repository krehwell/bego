package main

import (
    "fmt"
    "time"
)

func SwitchStatement() {
    i := "A"

    switch i {
    case "A" :
        fmt.Println("i is capital A")
        break;
    default:
        fmt.Println("i is another alphabet than A")
    }

    fmt.Println("----------")

    switch time.Now().Weekday() {
    case time.Saturday, time.Thursday:
        fmt.Println("today is Thursday")
        break;
    default:
        fmt.Println("today is not Thursday nor Saturday")
    }

    fmt.Println("----------")

    // switch as generic if/else statement
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    fmt.Println("----------")

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a boolean")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know what type it is, it's a %T\n", t)
        }
    }

    whatAmI(true)
    whatAmI(2304)
    whatAmI("Hey")
}
