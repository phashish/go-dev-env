package main

import "fmt"

func main() {
	favs := map[string][]string {
		"Vihaan": []string{"Books", "Guns", "Toys", "Cars"},
	}
	favs["Bond"] = []string{"guns", "girls", "cars"}
	favs["Disawal"] = []string{"go", "gum", "glue"}
	for key, val := range favs {
		for i, v := range val {
			fmt.Println("Name: ", key, " Fav item ", i+1, " is: ", v)
		}
	}

	if _, ok := favs["Bond"]; ok {
		delete(favs, "Bond")
		fmt.Println("We have deleted Bond's favs")
	} else {
		fmt.Println("We do not have Bond's favs")
	}

	if _, ok := favs["Divya"]; ok {
		delete(favs, "Divya")
		fmt.Println("We have deleted Divya's favs")
	} else {
		fmt.Println("We do not have Divya's favs")
	}

	for key, val := range favs {
		for i, v := range val {
			fmt.Println("Name: ", key, " Fav item ", i+1, " is: ", v)
		}
	}

}
