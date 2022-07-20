package main

import "fmt"

func main() {
	favs := map[string][]string {}
	favs["Bond"] = []string{"guns", "girls", "cars"}
	favs["Disawal"] = []string{"go", "gum", "glue"}
	fmt.Println(favs)
	for key, val := range favs {
		for i, v := range val {
			fmt.Println("Name: ", key, " Fav item ", i+1, " is: ", v)
		}
	}
}
