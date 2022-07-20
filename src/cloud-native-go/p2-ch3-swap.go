package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("3", "4")
	fmt.Println(a, b)
}
