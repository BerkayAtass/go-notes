package main

import "fmt"

func main() {
	temp := "one"
	{
		temp := "two"
		fmt.Println(temp)
	}
	fmt.Println(temp)
}
