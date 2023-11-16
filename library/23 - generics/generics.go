package main

import "fmt"

func main() {

	var number int = 1
	var number2 float64 = 5.8

	fmt.Println(addOne(number))
	fmt.Println(addOne(number2))

}

func addOne[n int | float64](number n) n {
	return number + 1
}
