package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func1(&wg)
	go func2(&wg)

	fmt.Println("Waiting for goroutines to finish")
	wg.Wait()
	fmt.Println("\nGoroutines finished")
}

func func1(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Func1 Done")
	wg.Done()
}

func func2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("Func2 Done")
}
