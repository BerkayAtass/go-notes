package main

import "fmt"

type group struct {
	name         string
	memberNumber int
}

func (g group) print() {
	fmt.Println("Group name : ", g.name, ", and member count : ", g.memberNumber)
}

func main() {

	var p1 group
	p1.name = "Yavuzlar"
	p1.memberNumber = 40

	p1.print()

}
