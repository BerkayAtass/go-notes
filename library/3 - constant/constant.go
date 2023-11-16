package main

import "fmt"

func main() {
	const name string = "SIBER VATAN"
	const group string = "YAVUZLAR"
	const field = "back-end"

	fmt.Println(name, "-", group, "-", field)

	//contant can not be changed
	//field = "front-end"

	fmt.Println(name, "-", group, "-", field)
}
