package red

import (
	"fmt"
	"time"
)

func Menu_marchand(player1 Character) Character {
	var i int
	Printfct("Bienvenu chez le marchand, que voulez vous acheter :\n\n1: Potion de vie (2 Pieces d'or) \n2: Potion de poison (2 Pieces d'or)\n3: Agrandir l'inventaire (30 Pieces d'Or)\n4: Peau de Troll (5 Pieces d'Or)\n5: Plume de Corbeau (5 Pieces d'Or)\n6: Cuir de Sanglier (5 Pieces d'Or)\n7: Fourrure de loup (5 Pieces d'Or)\n\n0: Retour", 30, 8)
	for {
		fmt.Scanln(&i)
		switch i {
		case 0:
			return player1
		case 1:
			player1 = Acheter_potion_vie(player1)
			Printfct("Bienvenu chez le marchand, que voulez vous acheter :\n\n1: Potion de vie (Gratuit) \n2: Potion de poison (2 Pieces d'or)\n3: Agrandir l'inventaire (30 Pieces d'Or)\n4: Peau de Troll (5 Pieces d'Or)\n5: Plume de Corbeau (5 Pieces d'Or)\n6: Cuir de Sanglier (5 Pieces d'Or)\n7: Fourrure de loup (5 Pieces d'Or)\n\n0: Retour", 1, 10)
		case 2:
      player1 = Acheter_potion_poison(player1)
			Printfct("Bienvenu chez le marchand, que voulez vous acheter :\n\n1: Potion de vie (Gratuit) \n2: Potion de poison (2 Pieces d'or)\n3: Agrandir l'inventaire (30 Pieces d'Or)\n4: Peau de Troll (5 Pieces d'Or)\n5: Plume de Corbeau (5 Pieces d'Or)\n6: Cuir de Sanglier (5 Pieces d'Or)\n7: Fourrure de loup (5 Pieces d'Or)\n\n0: Retour", 1, 10)
		case 3:
			player1 = UpgradeInventorySlot(player1)
			Printfct("Bienvenu chez le marchand, que voulez vous acheter :\n\n1: Potion de vie (Gratuit) \n2: Potion de poison (2 Pieces d'or)\n3: Agrandir l'inventaire (30 Pieces d'Or)\n4: Peau de Troll (5 Pieces d'Or)\n5: Plume de Corbeau (5 Pieces d'Or)\n6: Cuir de Sanglier (5 Pieces d'Or)\n7: Fourrure de loup (5 Pieces d'Or)\n\n0: Retour", 1, 10)
		case 4:
			player1 = Acheter_Peau_de_Troll(player1)
			Printfct("Bienvenu chez le marchand, que voulez vous acheter :\n\n1: Potion de vie (Gratuit) \n2: Potion de poison (2 Pieces d'or)\n3: Agrandir l'inventaire (30 Pieces d'Or)\n4: Peau de Troll (5 Pieces d'Or)\n5: Plume de Corbeau (5 Pieces d'Or)\n6: Cuir de Sanglier (5 Pieces d'Or)\n7: Fourrure de loup (5 Pieces d'Or)\n\n0: Retour", 1, 10)
		case 5:
			player1 = Acheter_Plume_de_Corbeau(player1)
			Printfct("Bienvenu chez le marchand, que voulez vous acheter :\n\n1: Potion de vie (Gratuit) \n2: Potion de poison (2 Pieces d'or)\n3: Agrandir l'inventaire (30 Pieces d'Or)\n4: Peau de Troll (5 Pieces d'Or)\n5: Plume de Corbeau (5 Pieces d'Or)\n6: Cuir de Sanglier (5 Pieces d'Or)\n7: Fourrure de loup (5 Pieces d'Or)\n\n0: Retour", 1, 10)
		case 6:
			player1 = Acheter_Cuir_de_Sanglier(player1)
			Printfct("Bienvenu chez le marchand, que voulez vous acheter :\n\n1: Potion de vie (Gratuit) \n2: Potion de poison (2 Pieces d'or)\n3: Agrandir l'inventaire (30 Pieces d'Or)\n4: Peau de Troll (5 Pieces d'Or)\n5: Plume de Corbeau (5 Pieces d'Or)\n6: Cuir de Sanglier (5 Pieces d'Or)\n7: Fourrure de loup (5 Pieces d'Or)\n\n0: Retour", 1, 10)
		case 7:
			player1 = Acheter_Fourrure_de_loup(player1)
			Printfct("Bienvenu chez le marchand, que voulez vous acheter :\n\n1: Potion de vie (Gratuit) \n2: Potion de poison (2 Pieces d'or)\n3: Agrandir l'inventaire (30 Pieces d'Or)\n4: Peau de Troll (5 Pieces d'Or)\n5: Plume de Corbeau (5 Pieces d'Or)\n6: Cuir de Sanglier (5 Pieces d'Or)\n7: Fourrure de loup (5 Pieces d'Or)\n\n0: Retour", 1, 10)
		default:
			println("invalid")
		}
	}
}

func Acheter_potion_vie(player1 Character) Character {
	if !TotalItems(player1.inventaire) {
		if player1.inventaire.piece_or >= 2 {
			player1.inventaire.potion += 1
			player1.inventaire.piece_or -= 2
			println("Achat en cours : Potion de vie")
			time.Sleep(3 * time.Second)
			println("Achat effectué : Vous venez d'obtenir une potion de vie")
		} else {
			fmt.Println("Ton inventaire est plein ! Impossible d’ajouter plus d’items.")
		}
	} else {
		fmt.Println("Vous n'avez pas assez de pièces d'or.")
	}
	return player1
}

func Acheter_potion_poison(player1 Character) Character {
	if !TotalItems(player1.inventaire) {
		if player1.inventaire.piece_or >= 2 {
			player1.inventaire.potion_poison += 1
			player1.inventaire.piece_or -= 2
			println("Achat en cours : Potion de poison")
			time.Sleep(3 * time.Second)
			println("Achat effectué : Vous venez d'obtenir une potion de poison")
		} else {
			fmt.Println("Ton inventaire est plein ! Impossible d'ajouter plus d'items.")
		}
	} else {
		fmt.Println("Vous n'avez pas assez de pièces d'or.")
	}
	return player1
}



func Acheter_Peau_de_Troll(player1 Character) Character {
	if !TotalItems(player1.inventaire) {
		if player1.inventaire.piece_or >= 2 {
			player1.inventaire.Peau_de_Troll += 1
			player1.inventaire.piece_or -= 5
			println("Achat en cours : Peau de Troll")
			time.Sleep(3 * time.Second)
			println("Achat effectué : Vous venez d'obtenir une Peau de Troll")
		} else {
			fmt.Println("Ton inventaire est plein ! Impossible d'ajouter plus d'items.")
		}
	} else {
		fmt.Println("Vous n'avez pas assez de pièces d'or.")
	}
	return player1
}

func Acheter_Plume_de_Corbeau(player1 Character) Character {
	if !TotalItems(player1.inventaire) {
		if player1.inventaire.piece_or >= 2 {
			player1.inventaire.Plume_de_Corbeau += 1
			player1.inventaire.piece_or -= 5
			println("Achat en cours : Plume de Corbeau")
			time.Sleep(3 * time.Second)
			println("Achat effectué : Vous venez d'obtenir une Plume de Corbeau")
		} else {
			fmt.Println("Ton inventaire est plein ! Impossible d'ajouter plus d'items.")
		}
	} else {
		fmt.Println("Vous n'avez pas assez de pièces d'or.")
	}
	return player1
}

func Acheter_Cuir_de_Sanglier(player1 Character) Character {
	if !TotalItems(player1.inventaire) {
			if player1.inventaire.piece_or >= 2 {
			player1.inventaire.Cuir_de_Sanglier += 1
			player1.inventaire.piece_or -= 5
			println("Achat en cours : Cuir de Sanglier")
			time.Sleep(3 * time.Second)
			println("Achat effectué : Vous venez d'obtenir un Cuir de Sanglier")
		} else {
			fmt.Println("Ton inventaire est plein ! Impossible d'ajouter plus d'items.")
		}
	} else {
		fmt.Println("Vous n'avez pas assez de pièces d'or.")
	}
	return player1
}

func Acheter_Fourrure_de_loup(player1 Character) Character {
	if !TotalItems(player1.inventaire) {
		player1.inventaire.Fourrure_de_loup += 1
		player1.inventaire.piece_or -= 5
		println("Achat en cours : Fourrure de loup")
		time.Sleep(3 * time.Second)
		println("Achat effectué : Vous venez d'obtenir une Fourrure de loup")
	} else {
		fmt.Println("Ton inventaire est plein ! Impossible d'ajouter plus d'items.")
	}
		return player1
}


