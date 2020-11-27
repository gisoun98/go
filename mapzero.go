package main

import (
	"fmt"
)

func main() {
	numbers := make(map[string]int)
	numbers["I've been assigned"] = 2
	numbers["I've been assigned"]++
	numbers["I've been assigned"]++
	
	//fmt.Printf("%d\n", numbers["I've been assigned"]++)
	fmt.Printf(numbers["I've been assigned"]++)
	fmt.Printf("%#v\n", numbers["I have not been assigned"])
}
