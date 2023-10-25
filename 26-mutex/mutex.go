package main

import (
	"fmt"
	"sync"
)

var mt sync.Mutex

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	var balance = 100
	fmt.Println("Initial balance:", balance)

	go addMoney(&wg, &balance, 10)
	go reduceMoney(&wg, &balance, 20)

	wg.Wait()
}

func addMoney(wg *sync.WaitGroup, balance *int, amount int) {
	mt.Lock()
	*balance += amount
	fmt.Println("Balance after adding:", *balance)
	mt.Unlock()
	wg.Done()
}

func reduceMoney(wg *sync.WaitGroup, balance *int, amount int) {
	mt.Lock()
	*balance -= amount
	fmt.Println("Balance after reducing:", *balance)
	mt.Unlock()
	wg.Done()
}
