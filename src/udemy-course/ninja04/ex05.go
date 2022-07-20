package main

import "fmt"

func main() {
	x := []int{33, 44, 55, 66, 77, 88, 99}
	fmt.Println(x)
    i := 4
	x = append(x[:i], x[i+2:]...)
	fmt.Println(x)

}
