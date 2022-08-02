package main

import "fmt"

func main() {
	x := struct{
		nameNumber map[string]int
		address []string
	}{
		nameNumber: map[string]int{
			"Ashish": 7899916782,
		},
		address: []string{"G601", "Sector 4", "Aaryavart"},
	}

	fmt.Println(x)
}
