package red

import "fmt"

type Inventory struct {
	baton int
	pomme int
	armure_cuivre int
	gourde int
}

func initInventory() Inventory {
	var inventaire_player1 Inventory
	inventaire_player1.baton = 1
	inventaire_player1.pomme = 0
	inventaire_player1.armure_cuivre = 0
	inventaire_player1.gourde = 0
	return inventaire_player1
}


func AccessInventory() {
	inventaire1_player1 := initInventory()
	fmt.Printf(" Voici votre inventaire :\n\n Baton = %v\n Pomme = %v\n Armure Cuivre = %v\n Gourde = %v\n", inventaire1_player1.baton, inventaire1_player1.pomme, inventaire1_player1.armure_cuivre, inventaire1_player1.gourde)
}
