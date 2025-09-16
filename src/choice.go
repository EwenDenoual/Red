package red

import "fmt"

func Menu(chara Character) Character {
	var i int
	Printfct("1: Character stats\n2: Inventory\n\n0: Back", 15, 3)
	for {
		fmt.Scanln(&i)
		switch i {
		case 0:
			return chara
		case 1:
			DisplayInfo(chara)
			Printfct("1: Character stats\n2: Inventory\n\n0: Back", 0, 9)
		case 2:
			DisplayInventory(chara)
			Printfct("1: Character stats\n2: Inventory\n\n0: Back", 2, 7)
		// case 3:
		// 	println("la tour")
		default:
			println("invalid")
		}
	}
}

func Choice(chara Character) {
	var i int
	Printfct("Welcome to the game :\n\n1: Acces menu\n\n0: Exit", 15, 4)
	for {
		fmt.Scanln(&i)
		switch i {
		case 0:
			return
		case 1:
			chara = Menu(chara)
			Printfct("Welcome to the game :\n\n1: Acces menu\n\n0: Exit", 15, 4)
		// case 2:
		// 	println("la foret")
		// case 3:
		// 	println("la tour")
		default:
			println("invalid")
		}
	}
}
