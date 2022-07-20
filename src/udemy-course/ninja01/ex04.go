package main

import "fmt"

type ash int
var x ash = 42

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
