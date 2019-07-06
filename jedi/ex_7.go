package jedi

import "fmt"

// Ex7 tdb
func Ex7() {
	x := 43
	if x < 42 {
		fmt.Println("я обосрался")
	} else if x > 44 {
		fmt.Println("я не обосрался")
	} else {
		fmt.Println("я опять обосрался")
	}
}
