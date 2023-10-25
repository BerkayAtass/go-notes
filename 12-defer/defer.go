package main

import "fmt"

func main() {
	defer fmt.Println("this is first")
	defer fmt.Println("this is second")
	defer fmt.Println("this is third")
	fmt.Println("this is fourth")
}
