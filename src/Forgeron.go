package red

import (
	"fmt"
	"time"
)

func Menu_forgeron(player1 Character) Character {
	var i int
	Printfct("Bienvenu dans chez le forgeron, que voulez vous fabriquer : \n\n1: Chapeau de l’aventurier (1 Plume de Corbeau et 1 Cuir de Sanglier)\n2: Tunique de l’aventurier (2 Fourrure de loup et 1 Peau de Troll)\n3: Bottes de l’aventurier (1 Fourrure de loup et 1 Cuir de Sanglier)\n\n0: Exit", 30, 5)
	for {
		fmt.Scanln(&i)
		switch i {
		case 0:
			return player1
		case 1:
			Chapeau_de_l_aventurier(player1)
			Printfct("Bienvenu dans chez le forgeron, que voulez vous fabriquer : \n\n1: Chapeau de l’aventurier (1 Plume de Corbeau et 1 Cuir de Sanglier)\n2: Tunique de l’aventurier (2 Fourrure de loup et 1 Peau de Troll)\n3: Bottes de l’aventurier (1 Fourrure de loup et 1 Cuir de Sanglier)\n\n0: Exit", 1, 7)
		case 2:
			Tunique_de_l_aventurier(player1)
			Printfct("Bienvenu dans chez le forgeron, que voulez vous fabriquer : \n\n1: Chapeau de l’aventurier (1 Plume de Corbeau et 1 Cuir de Sanglier)\n2: Tunique de l’aventurier (2 Fourrure de loup et 1 Peau de Troll)\n3: Bottes de l’aventurier (1 Fourrure de loup et 1 Cuir de Sanglier)\n\n0: Exit", 1, 7)
		case 3:
			Bottes_de_l_aventurier(player1)
			Printfct("Bienvenu dans chez le forgeron, que voulez vous fabriquer : \n\n1: Chapeau de l’aventurier (1 Plume de Corbeau et 1 Cuir de Sanglier)\n2: Tunique de l’aventurier (2 Fourrure de loup et 1 Peau de Troll)\n3: Bottes de l’aventurier (1 Fourrure de loup et 1 Cuir de Sanglier)\n\n0: Exit", 1, 7)
		default:
			println("invalid")
		}
	}
}

func Chapeau_de_l_aventurier(player1 Character) Character {
	
	if player1.inventaire.piece_or >= 5 {
		player1.inventaire.piece_or -= 5
		if player1.inventaire.Plume_de_Corbeau > 0 || player1.inventaire.Cuir_de_Sanglier > 0 {
			println("Votre Chapeau de l'aventurier est en cours de fabrication ")
			time.Sleep(3 *time.Second)
			println("Votre Chapeau de l'aventurier a ete fabriqué ")
			player1.inventaire.Plume_de_Corbeau -= 1
			player1.inventaire.Cuir_de_Sanglier -= 1
			player1.inventaire.Chapeau_de_l_aventurier += 1
		} else {
			println("Vous n'avez pas assez de ressources")
		}
	} else {
		println("Vous n'avez pas assez de pieces")
	}
	return player1
}

func Tunique_de_l_aventurier(player1 Character) Character {
	
	if player1.inventaire.piece_or >= 5 {
		player1.inventaire.piece_or -= 5
		if player1.inventaire.Fourrure_de_loup == 2 || player1.inventaire.Peau_de_Troll == 1 {
			println("Votre Tunique de l'aventurier est en cours de fabrication ")
			time.Sleep(3 *time.Second)
			println("Votre Tunique de l'aventurier a ete fabriqué ")
			player1.inventaire.Fourrure_de_loup -= 2
			player1.inventaire.Peau_de_Troll -= 1
			player1.inventaire.Tunique_de_l_aventurier += 1
		} else {
			println("Vous n'avez pas assez de ressources")
		}
	} else {
		println("Vous n'avez pas assez de pieces")
	}
	return player1
}

func Bottes_de_l_aventurier(player1 Character) Character {
	
	if player1.inventaire.piece_or >= 5 {
		player1.inventaire.piece_or -= 5
		if player1.inventaire.Fourrure_de_loup > 1 || player1.inventaire.Cuir_de_Sanglier > 1 {
			println("Vos Bottes de l'aventurier sont en cours de fabrication ")
			time.Sleep(3 *time.Second)
			println("Vos Bottes de l'aventurier ont ete fabriqué ")
			player1.inventaire.Fourrure_de_loup -= 1
			player1.inventaire.Cuir_de_Sanglier -= 1
			player1.inventaire.Bottes_de_l_aventurier += 1
		} else {
			println("Vous n'avez pas assez de ressources")
		}
	} else {
		println("Vous n'avez pas assez de pieces")
	}
	return player1
}
