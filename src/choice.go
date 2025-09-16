package red

import "fmt"

func Menu(chara Character) Character {
	var i int
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n1: Character stats\n2: Inventory\n\n0: Back\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n			")
	for {
		fmt.Scanln(&i)
		switch i {
		case 0:
			return chara
		case 1:
			DisplayInfo(chara)
			fmt.Println("\n\n1: Character stats\n2: Inventory\n\n0: Back\n\n\n\n\n\n\n			")
		case 2:
		 	DisplayInventory(chara)
			fmt.Println("\n\n1: Character stats\n2: Inventory\n\n0: Back\n\n\n\n\n\n\n\n\n\n\n			")
		// case 3:
		// 	println("la tour")
		default:
			println("invalid")
		}
	}
}

func Choice(chara Character) {
	var i int
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nWelcome to the game :\n\n1: Acces menu\n\n0: Exit\n\n\n\n\n\n\n\n\n\n\n\n\n\n			")
	for {
		fmt.Scanln(&i)
		switch i {
		case 0:
			return
		case 1:
			chara = Menu(chara)
			fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nWelcome to the game :\n\n1: Acces menu\n\n0: Exit\n\n\n\n\n\n\n\n\n\n\n\n\n\n			")
		// case 2:
		// 	println("la foret")
		// case 3:
		// 	println("la tour")
		default:
			println("invalid")
		}
	}
}
