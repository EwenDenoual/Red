package red

import (
	"fmt"
)

func Printfct(str string, before, nb int) {
	for range before {
		fmt.Printf("\n")
	}
	fmt.Println(str)
	for range 18 - nb {
		fmt.Printf("\n")
	}
}
