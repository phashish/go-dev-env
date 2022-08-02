package main

import "fmt"

func main() {
	type person struct{
		first string
		last string
		fav_ice_creams []string
	}

	ashish := person{
		first: "Ashish",
		last: "Disawal",
		fav_ice_creams: []string{"vanilla", "chocolate"},
	}

	divya := person{first: "Divya"}
	divya.last = "Disawal"
	divya.fav_ice_creams = []string{"strawberry", "gavava"}
	fmt.Println(ashish.fav_ice_creams)
	fmt.Println(divya)
}
