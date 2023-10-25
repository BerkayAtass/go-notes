package main

import (
	"fmt"
	"time"
)

func main() {

	//create a channel
	k := make(chan bool)

	//create a goroutine
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Hello")
		k <- true
	}()

	//wait for the goroutine to finish
	<-k

	fmt.Println("Done")

}
