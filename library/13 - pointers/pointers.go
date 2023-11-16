package main

import (
	"fmt"
	"time"
)

func main() {

	temp := 5
	addDefault(temp)
	fmt.Println(temp)

	time.Sleep(1 * time.Second)

	addPointer(&temp)
	fmt.Println(temp)
}

func addDefault(num int) {
	num += 10
}

func addPointer(num *int) {
	*num += 10
}
