package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("can't work with 42")
    }

    return arg + 3, nil
}

type myOwnCustomErrorStruct struct {
    arg int
    prob string
}

func (e *myOwnCustomErrorStruct) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {
        return -1, &myOwnCustomErrorStruct{arg: arg, prob: "can't work with it"}
    }

    return arg + 3, nil
}

func Errors() {
    for _, i := range []int{7, 42} {
        if num, err := f1(i); err != nil {
            fmt.Println("f1 failed:", err)
        } else {
            fmt.Println("f1 worked:", num)
        }
    }

    for _, i := range []int{7, 42} {
        if num, err := f2(i); err != nil {
            fmt.Println("f2 failed:", err)
        } else {
            fmt.Println("f2 worked:", num)
        }
    }

    _, e := f2(42)
    // `e.(*argError)` is type assertion / type casting in golang
    if ae, ok := e.(*myOwnCustomErrorStruct); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
