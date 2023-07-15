package main

import "fmt"

func main(){
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Err input num: %v", err)
		return
	}
	switch num {
	case 0:
		fmt.Println("OK")
	case 1, 2, 3:
		fmt.Println("Oh, small Integers!")
	case 4:
	case 5:
		fmt.Println("It's either 4 or 5")
		fallthrough
	case 6:
		fmt.Println("Guess what you'll get")
	default:
		fmt.Println("This is the default branch")
	}
}