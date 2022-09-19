package main

import (
	"fmt"
	"time"
)

// basically it's a setInterval in js
func Tickers() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case time := <-ticker.C:
				fmt.Println("Timer at", time)
			case <-done:
				return
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true

	fmt.Println("Ticker stopped")
}
