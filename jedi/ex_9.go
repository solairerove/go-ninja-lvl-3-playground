package jedi

import (
	"fmt"
)

// Ex9 tbd
func Ex9() {

	favSport := "tennis"

	switch favSport {
	case "tennis":
		fmt.Println("tennis")

	case "pool":
		fmt.Println("pool")

	default:
		fmt.Println("default")
	}
}
