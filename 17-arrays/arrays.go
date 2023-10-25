package main

import "fmt"

func main() {
	var temp [3]string
	temp[0] = "Yavuzlar"
	temp[1] = "Siber"
	temp[2] = "Vatan"

	fmt.Println(temp)

	fmt.Println("--------------------")

	var temp2 [3]string = [3]string{"Yavuzlar", "Siber", "Vatan"}

	fmt.Println(temp2)

	fmt.Println("--------------------")

	var temp3 = []string{}
	temp3 = append(temp3, "Yavuzlar")
	temp3 = append(temp3, "Siber")
	temp3 = append(temp3, "Vatan")

	fmt.Println(temp3)

	fmt.Println("--------------------")

}
