package main

import "fmt"

type group struct {
	name1, name2, name3 string
}

func main() {
	//var m map[string]group
	//m = make(map[string]group)

	//copilot shorted version of doing the same thing
	m := make(map[string]group)
	m["group1"] = group{"SiberVatan", "Yavuzlar", "Web Security"}
	m["group2"] = group{"Zayotem", "Cuberium", "Blockchain"}
	fmt.Println(m["group1"])
	fmt.Println(m["group2"])
}
