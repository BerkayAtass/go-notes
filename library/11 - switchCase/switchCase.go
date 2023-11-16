package main

import "fmt"

func main() {

	a := 10

	switch {
	case a == 10:
		fmt.Println("a is 10")
	case a < 20:
		fmt.Println("a is less than 20")
	}

	i := 15

	switch {
	case i == 15:
		fmt.Println("i is 15")
		fallthrough
	case i < 20:
		fmt.Println("i is less than 20")
	}

}
