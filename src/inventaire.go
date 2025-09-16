package red

import (
	"fmt"
	"time"
)

type Inventory struct {
	potion int
}

func initInventory() Inventory {
	var inventaire_player1 Inventory
	inventaire_player1.potion = 3
	return inventaire_player1
}


func DisplayInventory(player1 Character) {
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nInventaire : ")
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
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nVous disposez de %v pv sur %v pv\n", player1.pv, player1.pv_max)
		
	} else {
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nVous ne disposez pas de potion dans votre inventaire.\n")
	}
	
}

func PoisonPot(player1 Character) {
	player1.pv -= 10
	fmt.Printf("Vous venez de recevoir une potion de poison, voici vos PV : %v sur %v \n", player1.pv, player1.pv_max)

	time.Sleep(1 *time.Second)
	player1.pv -= 10
	fmt.Printf("La potion de poison fait effet, voici vos PV : %v sur %v \n", player1.pv, player1.pv_max)

	time.Sleep(1 *time.Second)
	player1.pv -= 10
	fmt.Printf("Vous perdez de plus en plus de vie, voici vos PV : %v sur %v \n", player1.pv, player1.pv_max)

	time.Sleep(1 *time.Second)
	fmt.Printf("La potion a fait effet, voici vos PV : %v sur %v \n", player1.pv, player1.pv_max)
}
