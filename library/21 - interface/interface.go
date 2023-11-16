package main

import "fmt"

type group interface {
	groupName()
}

type Yavuzlar struct {
}

func (y Yavuzlar) groupName() {
	fmt.Println("Yavuzlar")
}

type Zayotem struct {
}

func (z Zayotem) groupName() {
	fmt.Println("Zayotem")
}

func main() {
	var g group
	g = Yavuzlar{}
	g.groupName()
	g = Zayotem{}
	g.groupName()
}
