package red

import (
	"fmt"
	"time"
)

type Inventory struct {
	potion        int
	potion_poison int
	piece_or      int

	Plume_de_Corbeau int
	Cuir_de_Sanglier int
	Fourrure_de_loup int
	Peau_de_Troll    int

	Chapeau_de_l_aventurier int
	Tunique_de_l_aventurier int
	Bottes_de_l_aventurier  int

	size_max int
}

func TotalItems(inv Inventory) bool {
	tot := inv.potion +
		inv.potion_poison +
		inv.Plume_de_Corbeau +
		inv.Cuir_de_Sanglier +
		inv.Fourrure_de_loup +
		inv.Peau_de_Troll +
		inv.Chapeau_de_l_aventurier +
		inv.Tunique_de_l_aventurier +
		inv.Bottes_de_l_aventurier
	return tot >= inv.size_max
}

func initInventory() Inventory {
	var inventaire_player1 Inventory
	inventaire_player1.potion = 3
	inventaire_player1.potion_poison = 0

	inventaire_player1.piece_or = 10

	inventaire_player1.Plume_de_Corbeau = 0
	inventaire_player1.Cuir_de_Sanglier = 0
	inventaire_player1.Fourrure_de_loup = 0
	inventaire_player1.Peau_de_Troll = 0

	inventaire_player1.Chapeau_de_l_aventurier = 0
	inventaire_player1.Tunique_de_l_aventurier = 0
	inventaire_player1.Bottes_de_l_aventurier = 0

	inventaire_player1.size_max = 10

	return inventaire_player1
}

func DisplayInventory(player1 Character) Character {
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nInventaire : ")
	fmt.Printf("\n---------------------------------------------------------------\nPotion = %v\nPiece Or = %v\nPlume De Corbeau = %v\nCuir De Sanglier = %v\nFourrure De Loup = %v\nPeau De Troll = %v\nChapeau De L'Aventurier = %v\nTunique De L'aventurier = %v\nBottes De L'Aventurier = %v\n---------------------------------------------------------------\n\n", player1.inventaire.potion, player1.inventaire.piece_or, player1.inventaire.Plume_de_Corbeau,player1.inventaire.Cuir_de_Sanglier, player1.inventaire.Fourrure_de_loup, player1.inventaire.Peau_de_Troll, player1.inventaire.Chapeau_de_l_aventurier, player1.inventaire.Tunique_de_l_aventurier, player1.inventaire.Bottes_de_l_aventurier)
	player1 = Equipement(player1)
	return player1
}

func TakePot(player1 Character) Character {

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

func PoisonPot(player1 Character) Character {
	player1.pv -= 10
	fmt.Printf("Vous venez de recevoir une potion de poison, voici vos PV : %v sur %v \n", player1.pv, player1.pv_max)

	time.Sleep(1 * time.Second)
	player1.pv -= 10
	fmt.Printf("La potion de poison fait effet, voici vos PV : %v sur %v \n", player1.pv, player1.pv_max)

	time.Sleep(1 * time.Second)
	player1.pv -= 10
	fmt.Printf("Vous perdez de plus en plus de vie, voici vos PV : %v sur %v \n", player1.pv, player1.pv_max)

	time.Sleep(1 * time.Second)
	fmt.Printf("La potion a fait effet, voici vos PV : %v sur %v \n", player1.pv, player1.pv_max)
	return player1
}

func UpgradeInventorySlot(player1 Character) Character {
	if !TotalItems(player1.inventaire) {
		if player1.inventaire.piece_or >= 5 {
			player1.inventaire.size_max += 10
			player1.inventaire.piece_or -= 30
			println("Achat en cours : Agrandissement de l'inventaire")
			time.Sleep(3 * time.Second)
			println("Achat effectué : Vous venez d'obtenir un Agrandissement de l'inventaire")
		} else {
			fmt.Println("Vous n'avez pas assez de pièces d'or.")
		}
	} else {
		fmt.Println("Ton inventaire est plein ! Impossible d'ajouter plus d'items.")
	}
	return player1
}

func Equipement(player1 Character) Character {
	var i int
		Printfct("Equipement : \n\n1: Equiper tete\n2: Equiper torse\n3: Equiper botte\n\n0: Retour", 1, 20)
		for {
			fmt.Scanln(&i)
			switch i {
			case 0:
				return player1
			case 1:
				if player1.equipment.equipement_tete > 0 {
					Printfct("Vous êtes déjà équipé d’un chapeau.", 30,20)
					Printfct("Equipement : \n\n1: Equiper tete\n2: Equiper torse\n3: Equiper botte\n\n0: Retour", 7, 20)
				} else {
					player1 = Equipe_tete(player1)
					Printfct("Equipement : \n\n1: Equiper tete\n2: Equiper torse\n3: Equiper botte\n\n0: Retour", 7, 20)
				}
				
			case 2:
				if player1.equipment.equipement_torse > 1 {
					Printfct("Vous êtes déjà équipé d’une tunique.", 30,20)
					Printfct("Equipement : \n\n1: Equiper tete\n2: Equiper torse\n3: Equiper botte\n\n0: Retour", 7, 20)
				} else {
					player1 = Equipe_torse(player1)
					Printfct("Equipement : \n\n1: Equiper tete\n2: Equiper torse\n3: Equiper botte\n\n0: Retour", 7, 20)
				}
			case 3:
				if player1.equipment.equipement_pieds > 1 {
					Printfct("Vous êtes déjà équipé de bottes.", 30,20)
					Printfct("Equipement : \n\n1: Equiper tete\n2: Equiper torse\n3: Equiper botte\n\n0: Retour", 7, 20)
				} else {
					player1 = Equipe_pieds(player1)
					Printfct("Equipement : \n\n1: Equiper tete\n2: Equiper torse\n3: Equiper botte\n\n0: Retour", 7, 20)
				}
			default:
				println("invalid")
			}
		}
}


func Equipe_tete(player1 Character) Character {
	if player1.inventaire.Chapeau_de_l_aventurier >= 1 {
		Printfct("Votre chapeau de l’aventurier a été équipé.", 30,20)
		player1.inventaire.Chapeau_de_l_aventurier -= 1
		player1.equipment.equipement_tete += 1
		player1.pv_max += 10
		player1.pv += 10
	} else {
		Printfct("Vous ne disposez pas d’équipement de tête.", 30,20)
	}
	return player1
}

func Equipe_torse(player1 Character) Character {
	if player1.inventaire.Tunique_de_l_aventurier >= 1 {
		Printfct("Votre tunique de l’aventurier a été équipé.", 30,20)
		player1.inventaire.Tunique_de_l_aventurier -= 1
		player1.equipment.equipement_torse += 1
		player1.pv_max += 25
		player1.pv += 25
	} else {
		Printfct("Vous ne disposez pas d’équipement de torse.", 30,20)
	}
	return player1
}

func Equipe_pieds(player1 Character) Character {
	if player1.inventaire.Bottes_de_l_aventurier >= 1 {
		Printfct("Vos bottes de l’aventurier a été équipé.", 30,20)
		player1.inventaire.Bottes_de_l_aventurier -= 1
		player1.equipment.equipement_pieds += 1
		player1.pv_max += 15
		player1.pv += 15	
	} else {
		Printfct("Vous ne disposez pas d’équipement de pieds.", 30,20)
	}
	return player1
}
