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
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Time at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)

	ticker.Stop()
	done <- true

	fmt.Println("Ticker stopped")
}
