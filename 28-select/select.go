package main

import (
	"fmt"
	"time"
)

func main() {
	temp1 := make(chan string)
	temp2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		temp1 <- "SIBER"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		temp2 <- "VATAN"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-temp1:
			fmt.Println("Received", msg1)
		case msg2 := <-temp2:
			fmt.Println("Received", msg2)
		}
	}
}
