package main

import (
	"fmt"
	"time"
)

// setTimeout in js
func Timers() {
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 Fired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 Fired")
	}()

	// time.Sleep(4 * time.Second)

	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 Stopped")
	}

	time.Sleep(2 * time.Second)
}
