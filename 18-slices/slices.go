package main

import "fmt"

func main() {
	temp := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(temp)
	var temp1 []int = temp[2:5]
	fmt.Println(temp1)

	var temp2 []int = temp[:5]
	fmt.Println(temp2)

	var temp3 []int = temp[2:]
	fmt.Println(temp3)

}
