package red

import (
	"fmt"
	"time"
)

func Menu_marchand(player1 Character) Character {
	var i int
	Printfct("Bienvenu chez le marchand, que voulez vous acheter :\n\n1: Potion de vie\n2: Potion de poison\n\n0: Exit", 30, 5)
	for {
		fmt.Scanln(&i)
		switch i {
		case 0:
			return player1
		case 1:
			player1 = Acheter_potion_vie(player1)
			Printfct("Bienvenu chez le marchand, que voulez vous acheter :\n\n1: Potion de vie\n2: Potion de poison\n\n0: Exit", 1, 7)
		case 2:
			player1 = Acheter_potion_poison(player1)
			Printfct("Bienvenu chez le marchand, que voulez vous acheter :\n\n1: Potion de vie\n2: Potion de poison\n\n0: Exit", 1, 7)
		//case 3:
		//fonction(player1)
		//Printfct("", 1, 7)
		default:
			println("invalid")
		}
	}
}

func Acheter_potion_vie(player1 Character) Character {
	if !TotalItems(player1.inventaire) {
		player1.inventaire.potion += 1
		println("Achat en cours : Potion de vie")
		time.Sleep(3 * time.Second)
		println("Achat effectué : Vous venez d'obtenir une potion de vie")
	} else {
		fmt.Println("Ton inventaire est plein ! Impossible d’ajouter plus d’items.")
	}
	return player1
}

func Acheter_potion_poison(player1 Character) Character {
	player1.inventaire.potion_poison += 1 {
	println("Achat en cours : Potion de poison")
	time.Sleep(3 * time.Second)
	println("Achat effectué : Vous venez d'obtenir une potion de poison")
} else {
	fmt.Println("Ton inventaire est plein ! Impossible d'ajouter plus d'items.")
}
	return player1
}
