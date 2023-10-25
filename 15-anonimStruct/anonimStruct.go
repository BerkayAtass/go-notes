package main

import "fmt"

func main() {

	group := struct {
		name         string
		memberNumber int
	}{"Yavuzlar", 40}

	fmt.Println(group)
	fmt.Println("--------------------")
	fmt.Println(group.name)
	fmt.Println(group.memberNumber)

}
