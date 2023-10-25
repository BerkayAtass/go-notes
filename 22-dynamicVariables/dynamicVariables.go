package main

import "fmt"

type dynamic interface{}

func main() {

	var temp dynamic

	temp = 1

	fmt.Println(temp)

	temp = "SiberVatan"

	fmt.Println(temp)

	//or
	//this one is new in Go and more useful

	var temp1 any

	temp1 = 1

	fmt.Println(temp1)

	temp1 = "SiberVatan"

	fmt.Println(temp1)

}
