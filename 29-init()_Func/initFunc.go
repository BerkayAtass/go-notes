package main

import "fmt"

/*When an application runs in Golang, the first function that
runs is usually the main() function. Sometimes, prerequisites
occur that we need to set at the start of the program. The init()
function offers us this opportunity. Let's add meaning to what I
wrote with a small example.
(https://go.kaanksc.com/boeluem-5/init-fonksiyonu-oen-yuekleme)*/

func init() {
	fmt.Println("Init func")
}

func main() {
	fmt.Println("Main func")
}
