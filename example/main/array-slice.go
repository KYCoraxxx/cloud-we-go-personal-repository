package main

import "fmt"

func main(){
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	var slice1 []int = arr[2:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("part1[%v] = %v\n", i, slice1[i])
	}

	slice2 := append(slice1, 0)
	slice2[0] = 0

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("part1[%v] = %v\n", i, slice1[i])
	}

	for i := 0; i < len(slice2); i++ {
		fmt.Printf("part2[%v] = %v\n", i, slice2[i])
	}
}