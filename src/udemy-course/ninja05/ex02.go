package main

import "fmt"

func main() {
	type vehicle struct{
		doors int
		color string
	}

	type truck struct{
		vehicle
		fourWheel bool
	}

	type sedan struct{
		vehicle
		luxury bool
	}

	v1 := vehicle{
		doors: 4,
		color: "pink",
	}

	fmt.Println(v1)

	t1 := truck{
		vehicle: vehicle{doors: 4, color: "grey"}, fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{doors: 2, color: "white"}, luxury: true,
	}
	
	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Println(t1.doors)
	fmt.Println(s1.doors)

}
