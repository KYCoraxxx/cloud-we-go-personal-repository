package main

import "fmt"

func main(){
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}
	i := 10
	for ; i > 0; i-- {
		fmt.Println("i = ", i)
	}
	i = 1024
	for i > 0 {
		fmt.Print(i % 10)
		i /= 10
	}

	for {
		fmt.Println("I will run forever")
		break
	}
}