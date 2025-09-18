package red

import (
	"fmt"
)

func Menu(chara Character) Character {
	var i int
	Printfct("1: Créer votre personnage\n2: Statistiques du personnage\n3: Inventaire\n4: Utiliser une potion\n\n0: Retour", 12, 3)
	for {
		fmt.Scanln(&i)
		switch i {
		case 0:
			return chara
		case 1 : 
			chara = CharacterCreation(chara)
			Printfct("1: Créer votre personnage\n2: Statistiques du personnage\n3: Inventaire\n4: Utiliser une potion\n\n0: Retour", 0, 3)
		case 2:
			DisplayInfo(chara)
			Printfct("1: Créer votre personnage\n2: Statistiques du personnage\n3: Inventaire\n4: Utiliser une potion\n\n0: Retour", 0, 11)
		case 3:
			if chara.nom != "" {
				chara = DisplayInventory(chara)
				Printfct("1: Créer votre personnage\n2: Statistiques du personnage\n3: Inventaire\n4: Utiliser une potion\n\n0: Retour", 20,3)
			} else {
				fmt.Printf("Veillez créer un personnage")
				Printfct("Bienvenue dans le jeu :\n\n1: Accéder au menu du personnage\n2: Aller chez le forgeron\n3: Aller chez le marchand\n4: Trouver un adversaire\n\n0: Retour", 2, 7)
			}
		case 4:
		 	chara = TakePot(chara)
			Printfct("1: Créer votre personnage\n2: Statistiques du personnage\n3: Inventaire\n4: Utiliser une potion\n\n0: Retour", 0, 10)
		default:
			println("invalid")
		}
	}
}

func Choice(chara Character) Character {
	var i int
  Printfct("Bienvenue dans le jeu :\n\n1: Accéder au menu du personnage\n2: Aller chez le forgeron\n3: Aller chez le marchand\n4: Trouver un adversaire\n\n0: Retour", 15, 5)
	for {
		fmt.Scanln(&i)
		switch i {
		case 0:
			return chara
		case 1:
			chara = Menu(chara)
			Printfct("Bienvenue dans le jeu :\n\n1: Accéder au menu du personnage\n2: Aller chez le forgeron\n3: Aller chez le marchand\n4: Trouver un adversaire\n\n0: Retour", 15, 5)
		case 2:
			if chara.nom != "" {
				chara = Menu_forgeron(chara)
				Printfct("Bienvenue dans le jeu :\n\n1: Accéder au menu du personnage\n2: Aller chez le forgeron\n3: Aller chez le marchand\n4: Trouver un adversaire\n\n0: Retour", 15, 5)
			} else {
				fmt.Printf("Veillez créer un personnage")
				Printfct("Bienvenue dans le jeu :\n\n1: Accéder au menu du personnage\n2: Aller chez le forgeron\n3: Aller chez le marchand\n4: Trouver un adversaire\n\n0: Retour", 2, 7)
			}
			case 3:
			if chara.nom != "" {
				chara = Menu_marchand(chara)
				Printfct("Bienvenue dans le jeu :\n\n1: Accéder au menu du personnage\n2: Aller chez le forgeron\n3: Aller chez le marchand\n4: Trouver un adversaire\n\n0: Retour", 15, 5)
			} else {
				fmt.Printf("Veillez créer un personnage")
				Printfct("Bienvenue dans le jeu :\n\n1: Accéder au menu du personnage\n2: Aller chez le forgeron\n3: Aller chez le marchand\n4: Trouver un adversaire\n\n0: Retour", 2, 7)
			}
			case 4:
			if chara.nom != "" {
		 		chara = fight(chara)
				Printfct("Bienvenue dans le jeu :\n\n1: Accéder au menu du personnage\n2: Aller chez le forgeron\n3: Aller chez le marchand\n4: Trouver un adversaire\n\n0: Retour", 15, 5)
			} else {
				fmt.Printf("Veillez créer un personnage")
				Printfct("Bienvenue dans le jeu :\n\n1: Accéder au menu du personnage\n2: Aller chez le forgeron\n3: Aller chez le marchand\n4: Trouver un adversaire\n\n0: Retour", 2, 7)

			}
		default:
			println("invalid")
		}
	}
}
