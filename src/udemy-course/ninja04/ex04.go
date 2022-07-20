package main

import "fmt"

func main() {
	x := []int{33, 44, 55}
	fmt.Println(x)
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58}
	x = append(x, y...)
	fmt.Println(x)
}
