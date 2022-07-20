package main

import "fmt"

func main() {
	x := []int{1, 22, 33, 44, 55, 66, 77, 88, 99, 10}
	fmt.Printf("x[:5]\t %v\n", x[:5])
	fmt.Printf("x[1:7]\t %v\n", x[1:7])
	fmt.Printf("x[5:]\t %v\n", x[5:])
	fmt.Printf("x[2:4]\t %v\n", x[2:4])
	fmt.Printf("x[8:]\t %v\n", x[8:])
	fmt.Printf("x[:]\t %v\n", x[:])
	fmt.Printf("%T\n", x)
}
