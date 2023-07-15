package main

import (
	"fmt"
	"sync"
)

func hello() {
	fmt.Println("Hello from goroutine")
}

func exampleMutex() {
	var m sync.Mutex
	m.Lock()
	go func() {
		// move m.Lock() here, the coroutine will be temperately put to pending, and the process finish, resource get recycled
		defer m.Unlock()
		fmt.Println("Hello from goroutine")
	}()

	m.Lock()
	defer m.Unlock()
	fmt.Println("finish")
}

func main() {
	// go hello()
	// time.Sleep(1 * time.Second)
	// fmt.Println("Hello from main")
	exampleMutex()
}