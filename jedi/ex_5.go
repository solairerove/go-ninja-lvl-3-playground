package jedi

import (
	"fmt"
)

// Ex5 tbd
func Ex5() {

	for i := 10; i <= 100; i++ {
		fmt.Printf("%v\t%v\n", i, i%4)
	}
}
