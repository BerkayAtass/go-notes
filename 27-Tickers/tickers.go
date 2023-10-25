package main

import (
	"fmt"
	"time"
)

func main() {
	repeat := time.NewTicker(1 * time.Second)
	finish := make(chan bool)

	go func() {
		for {
			select {
			case <-finish:
				return
			case time := <-repeat.C:
				fmt.Println("Time:", time)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	repeat.Stop()
	finish <- true
	fmt.Println("Ticker stopped")
}
