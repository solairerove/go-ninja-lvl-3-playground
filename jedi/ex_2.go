package jedi

import (
	"fmt"
)

// Ex2 printing from 1 to 10k
func Ex2() {

	for i := 65; i <= 90; i++ {
		for x := 0; x <= 2; x++ {
			fmt.Printf("%#U\n", i)
		}

	}
}
