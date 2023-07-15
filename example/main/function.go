package main

import (
	"fmt"
	"strconv"
)

func intToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func add(i, j int) int {
	return i + j
}

func sameInteger(i int, s string) (bool, error) {
	if si, err := strconv.Atoi(s); err != nil {
		return false, err
	} else {
		return si == i, nil
	}
}

func sum(x ...int) (result int) {
	for _, v := range x {
		result += v
	}
	return
}

func main(){
	var a [5]int
	for i := 0; i < len(a); i++ {
		a[i] = i;
	}
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum([]int{1, 2, 3}...))
	fmt.Println(sum(a[2:5]...)) 		// append ... to slice in parameter
}