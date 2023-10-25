package main

import "fmt"

func main() {

	var name, surname string

	fmt.Print("What is your name and surname? ")
	fmt.Scanf("%s %s", &name, &surname)

	fmt.Printf("Hello %s %s\n", name, surname)

}
