package red

import (
	"fmt"
)

type Inventory struct {
	potion int
}

func initInventory() Inventory {
	var inventaire_player1 Inventory
	inventaire_player1.potion = 3
	return inventaire_player1
}


func AccessInventory(player1 Character) {
	fmt.Printf("\n Potion = %v\n", player1.inventaire.potion)
}

func TakePot(player1 Character) {
	
	if player1.inventaire.potion > 0 {
		player1.pv += 50
		player1.inventaire.potion -= 1
		if player1.pv > player1.pv_max {
			player1.pv = player1.pv_max
		}
		DisplayInfo(player1)
		fmt.Printf("Vous disposez de %v pv sur %v pv\n", player1.pv, player1.pv_max)
		
	} else {
		fmt.Printf("Vous ne disposez pas de potion dans votre inventaire.\n")
	}
	
}
