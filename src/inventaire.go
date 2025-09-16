package red

import (
	"fmt"
	"time"
)

type Inventory struct {
	potion int
	potion_poison int
	piece_or int
	
	Plume_de_Corbeau int
	Cuir_de_Sanglier int
	Fourrure_de_loup int
	Peau_de_Troll int

	Chapeau_de_l_aventurier int
	Tunique_de_l_aventurier int
	Bottes_de_l_aventurier int
}

func initInventory() Inventory {
	var inventaire_player1 Inventory
	inventaire_player1.potion = 3

	inventaire_player1.piece_or = 10

	inventaire_player1.Plume_de_Corbeau = 2
	inventaire_player1.Cuir_de_Sanglier = 2
	inventaire_player1.Fourrure_de_loup = 2
	inventaire_player1.Peau_de_Troll = 2

	inventaire_player1.Chapeau_de_l_aventurier = 0
	inventaire_player1.Tunique_de_l_aventurier = 0
	inventaire_player1.Bottes_de_l_aventurier = 0




	return inventaire_player1
}


func DisplayInventory(player1 Character) {
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nInventaire : ")
	fmt.Printf("\n Potion = %v\n", player1.inventaire.potion)
}

func TakePot(player1 Character) Character{
	
	if player1.inventaire.potion > 0 {
		player1.pv += 50
		player1.inventaire.potion -= 1
		if player1.pv > player1.pv_max {
			player1.pv = player1.pv_max
		}
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nVous disposez de %v pv sur %v pv\n", player1.pv, player1.pv_max)
		
	} else {
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nVous ne disposez pas de potion dans votre inventaire.\n")
	}
	return player1
	
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
