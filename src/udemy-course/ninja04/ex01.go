package main

import "fmt"

func main() {
	x := [5]int{}
	for i, _ := range x {
		x[i] = i * i
		fmt.Printf("%v\n", x)
		fmt.Printf("%T\n", x)
	}
}
