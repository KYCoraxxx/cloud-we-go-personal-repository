package main

import (
	"fmt"
	"math/rand"
)

func GaussSum(n int) (result int64) {
	for i := 1; i <= n; i++ {
		result += int64(i)
	} 
	return result
}

func main() {
	//c := make(chan int64)
	//for i := 1; i < 101; i++ {
	//	go func() {
	//		c <- GaussSum(i)
	//	}()
	//	result := <- c			// the main must wait for something is in channel
	//	fmt.Println("result = ", result)
	//}
	//multi_length_chanel()
	selectCoroutine()
}

func multi_length_chanel() {
	c := make(chan int64, 3)	// dead lock if 2
	c <- 1
	c <- 2
	c <- 3
	for i := 0; i < 3; i++ {
		fmt.Println(<-c)
	}
}

func selectCoroutine() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		if rand.Intn(2) == 0 {
			c1 <- "c1"
		} else {
			c2 <- "c2"
		}
	}()
	for {
		select {
		case c := <- c1:
			fmt.Println(c)
			return
		case c := <- c2:
			fmt.Println(c)
			return
		default: // avoid blocking
		}
	}
}