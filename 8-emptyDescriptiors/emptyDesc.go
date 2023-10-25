package main

import "fmt"

func main() {

	process1, process2 := temp(12)

	fmt.Println(process1, process2)

	process3, _ := temp(16)
	_, process4 := temp(20)

	fmt.Println(process3, process4)

}

func temp(num int) (int, int) {
	process1 := num / 2
	process2 := num / 4
	return process1, process2
}
