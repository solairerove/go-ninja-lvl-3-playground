package jedi

import (
	"fmt"
)

// Ex4 tbd
func Ex4() {
	bornYear := 1996

	for {
		if bornYear > 2019 {
			break
		}

		fmt.Println(bornYear)
		bornYear++
	}
}
