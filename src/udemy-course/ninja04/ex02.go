package main

import "fmt"

func main() {
	x := []int{1, 22, 33, 44, 55, 66, 77, 88, 99, 10}
	for i, v := range x {
		fmt.Printf("%v\t%v\n", i, v)
	}
	fmt.Printf("%T\n", x)
}
