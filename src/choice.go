package red

import "fmt"

func TestChoice() {
	var i int
	fmt.Println("1 2 ou 3")
	for j := 0; j < 1; j++ {
		fmt.Scanln(&i)
		switch i {
		case 0:
			return
		case 1:
			println("la ville")
		case 2:
			println("la foret")
		case 3:
			println("la tour")
		default:
			println("invalid")
			j = -1
		}
	}
}
