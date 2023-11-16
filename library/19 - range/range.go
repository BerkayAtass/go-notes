package main

import "fmt"

var names []string = []string{"SiberVatan", "Yavuzlar", "Zayotem", "Cuberium"}

func main() {
	for a, b := range names {
		fmt.Printf("%d index: %s\n", a, b)
	}
}
