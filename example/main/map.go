package main

import "fmt"

func main(){
	m := map[string]int{
		"a" : 1,
		"b" : 2,
	}
	m["c"] = 3

	fmt.Println("m[a] = ", m["a"])

	if v, exist := m["d"]; exist {
		fmt.Println("m[d] = ", v)
	} else {
		fmt.Println("inexistent key in map")
	}

	delete(m, "a")

	for k, v := range m {
		fmt.Printf("m[%v] = %v\n", k, v)
	}
}