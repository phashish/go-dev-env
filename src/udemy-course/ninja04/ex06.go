package main

import "fmt"

func main() {
	states := make([]string, 0, 28)
	fmt.Printf("Cap: %v\tLen: %v\tStates: %v\n", cap(states), len(states), states)
	states = append(states, "Andhra Pradesh", "Arunachal Pradesh", "Assam", "Bihar", "Chhattisgarh", "Goa", "Gujrat", "Haryana", "Himachal Pradesh", "Jharkhand", "Karnataka", "Kerala", "Madhya Pradesh", "Maharashtra", "Manipur", "Meghalaya", "Mizoram", "Nagaland", "Odisha", "Punjab", "Rajasthan", "Sikkim", "Tamil Nadu", "Telangana", "Tripura", "Uttarakhand", "Uttar Pradesh", "West Bengal")
	fmt.Printf("Cap: %v\tLen: %v\tStates: %v\n", cap(states), len(states), states)
}
