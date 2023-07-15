package main

import (
	"fmt"
	"strings"
)

func main(){
	var s string
	fmt.Scanln(&s)
	if strings.HasPrefix(s, "yes") {
		fmt.Println("Cool.")
	} else if strings.HasPrefix(s, "no") {
		fmt.Println("Tell me why.")
	} else {
		fmt.Println("WTF?")
	}
}