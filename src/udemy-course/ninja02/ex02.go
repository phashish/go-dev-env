package main

import "fmt"

func main() {
	a := 7
	for a < 9999 {
		fmt.Printf("%d, %b, %#x\n", a, a, a)
		a = a << 1
	}	 
}
