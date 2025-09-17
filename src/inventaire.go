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
	fmt.Printf("\n\nPotion = %v\nPiece Or = %v\nPlume De Corbeau = %v\nCuir De Sanglier = %v\nFourrure De Loup = %v\nPeau De Troll = %v\nChapeau De L'Aventurier = %v\nTunique De L'aventurier = %v\nBottes De L'Aventurier = %v\n", player1.inventaire.potion, player1.inventaire.piece_or, player1.inventaire.Plume_de_Corbeau,player1.inventaire.Cuir_de_Sanglier, player1.inventaire.Fourrure_de_loup, player1.inventaire.Peau_de_Troll, player1.inventaire.Chapeau_de_l_aventurier, player1.inventaire.Tunique_de_l_aventurier, player1.inventaire.Bottes_de_l_aventurier)
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

// Fonction d'agrandissment de l'inventaire a ajouter quand la fonction "Limite d’inventaire" (Tache 12)

// func UpgradeInventorySlot(player1 Character) Character {
// 	var Inventory_max_use int = 3

// 	if player1.inventaire.piece_or > 30 {
// 		if Inventory_max_use > 0 {
// 			println("Vous venez d'agrandire votre inventaire de + 10 items")
// 			Inventory_max += 10
// 			Inventory_max_use -= 1
// 			player1.inventaire.piece_or -= 30

// 		}
// 	}
// 	return player1
// }
