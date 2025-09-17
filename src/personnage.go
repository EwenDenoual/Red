package red

import (
	"fmt"
	"time"
)

type Character struct {
	nom        string
	classe     string
	niveau     int
	pv_max     int
	pv         int
	st         stat
	inventaire Inventory
	spell      spellbook
	equipment Equipment
}


type Equipment struct {
	equipement_tete int
	equipement_torse int
	equipement_pieds int
}

func InitEquipment() Equipment {
	var equipement_player1 Equipment
	equipement_player1.equipement_tete = 0
	equipement_player1.equipement_torse = 0
	equipement_player1.equipement_pieds = 0
	return equipement_player1
}

func InitCharacter() Character {
	var player1 Character
	// player1.nom = " "
	player1.classe = ""
	player1.niveau = 1
	player1.pv_max = 100
	player1.pv = 40
	player1.st = initStat()
	player1.inventaire = initInventory()
	player1.spell = initSpellbook()
	return player1
}

func IsDead(player1 Character) bool {
	var i int
	if player1.pv == 0 {
		Printfct("Vous etes mort dommage !!!  :)", 30, 0)
		fmt.Scanln(&i)
		return true
	}
	return false
}

func DisplayInfo(player1 Character) {
	if player1.nom != "" {
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n---------------------------------------------------------------\n Nom : %v\n Classe : %v\n Niveau : %v\n Pv_max : %v\n Pv : %v\n---------------------------------------------------------------\n\n", player1.nom, player1.classe, player1.niveau, player1.pv_max, player1.pv)
	} else {
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\nVeuillez créer votre personnage\n\n")
	}
}

func CharacterCreation(player1 Character) Character {

	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nCréer / Modifier votre personnage : ")
	fmt.Println("Entrer le nom de votre personnage : ")
	fmt.Scan(&player1.nom)

	println("Entrer la classe de votre personnage : ")

	var i int
	Printfct("1: Classe Humain\n2: Classe Elfes\n3: Classe Nains\n\n0: Exit", 3,9)
	for {
		fmt.Scanln(&i)
		switch i {
		case 0: 
			return player1
		case 1 :
			player1.classe = "Humain"
			player1.pv_max = 100
			player1.pv = player1.pv_max/2
			Printfct("Classe Humain chose", 30,1)
			time.Sleep(2 *time.Second)
			return player1
		case 2 : 
			player1.classe = "Elfes"
			player1.pv_max = 80
			player1.pv = player1.pv_max/2
			Printfct("Classe Elfes chose", 30,1)
			time.Sleep(2 *time.Second)
			return player1
		case 3 :
			player1.classe = "Nains"
			player1.pv_max = 120
			player1.pv = player1.pv_max/2
			Printfct("Classe Nains chose", 30,1)
			time.Sleep(2 *time.Second)
			return player1
		}
	}
}