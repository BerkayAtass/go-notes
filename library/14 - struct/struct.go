package main

import "fmt"

func main() {

	type group struct {
		name         string
		memberNumber int
	}

	var p1 group
	p1.name = "Yavuzlar"
	p1.memberNumber = 40

	fmt.Println(p1)
	fmt.Println("--------------------")
	fmt.Println(p1.name)
	fmt.Println(p1.memberNumber)

}
