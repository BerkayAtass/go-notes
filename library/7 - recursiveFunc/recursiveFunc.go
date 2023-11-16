package main

import "fmt"

func main() {

	fmt.Println(recursive(5))

}

func recursive(num int) int {

	if num == 0 {
		return 1
	}
	return num * recursive(num-1)

}
