package red

import (
	"fmt"
)

func Menu(chara Character) Character {
	var i int
	Printfct("1: Créer votre personnage\n2: Character stats\n3: Inventory\n4: Use Potion\n\n0: Back", 12, 3)
	for {
		fmt.Scanln(&i)
		switch i {
		case 0:
			return chara
		case 1 : 
			chara = CharacterCreation(chara)
			Printfct("1: Créer votre personnage\n2: Character stats\n3: Inventory\n4: Use Potion\n\n0: Back", 0, 3)
		case 2:
			DisplayInfo(chara)
			Printfct("1: Créer votre personnage\n2: Character stats\n3: Inventory\n4: Use Potion\n\n0: Back", 0, 11)
		case 3:
			DisplayInventory(chara)
			Printfct("1: Créer votre personnage\n2: Character stats\n3: Inventory\n4: Use Potion\n\n0: Back", 0, 16)
		case 4:
		 	chara = TakePot(chara)
			Printfct("1: Créer votre personnage\n2: Character stats\n3: Inventory\n4: Use Potion\n\n0: Back", 0, 10)
		default:
			println("invalid")
		}
	}
}

func Choice(chara Character) Character {
	var i int
  Printfct("Welcome to the game :\n\n1: Acces menu\n2: Acces Forgeron\n3: Acces Marchand\n4: find an opponent\n\n0: Exit", 15, 5)
	for {
		fmt.Scanln(&i)
		switch i {
		case 0:
			return chara
		case 1:
			chara = Menu(chara)
			Printfct("Welcome to the game :\n\n1: Acces menu\n2: Acces Forgeron\n3: Acces Marchand\n4: find an opponent\n\n0: Exit", 15, 5)
		case 2:
			if chara.nom != "" {
				chara = Menu_forgeron(chara)
				Printfct("Welcome to the game :\n\n1: Acces menu\n2: Acces Forgeron\n3: Acces Marchand\n4: find an opponent\n\n0: Exit", 15, 5)
			} else {
				fmt.Printf("Veillez créer un personnage")
				Printfct("Welcome to the game :\n\n1: Acces menu\n2: Acces Forgeron\n3: Acces Marchand\n4: find an opponent\n\n0: Exit", 2, 7)
			}
			case 3:
			if chara.nom != "" {
				chara = Menu_marchand(chara)
				Printfct("Welcome to the game :\n\n1: Acces menu\n2: Acces Forgeron\n3: Acces Marchand\n4: find an opponent\n\n0: Exit", 15, 5)
			} else {
				fmt.Printf("Veillez créer un personnage")
				Printfct("Welcome to the game :\n\n1: Acces menu\n2: Acces Forgeron\n3: Acces Marchand\n4: find an opponent\n\n0: Exit", 2, 7)
			}
			case 4:
			if chara.nom != "" {
		 		chara = fight(chara)
				Printfct("Welcome to the game :\n\n1: Acces menu\n2: Acces Forgeron\n3: Acces Marchand\n4: find an opponent\n\n0: Exit", 15, 5)
			} else {
				fmt.Printf("Veillez créer un personnage")
				Printfct("Welcome to the game :\n\n1: Acces menu\n2: Acces Forgeron\n3: Acces Marchand\n4: find an opponent\n\n0: Exit", 2, 7)

			}
		default:
			println("invalid")
		}
	}
}
