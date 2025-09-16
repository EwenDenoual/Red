package red

import (
	"fmt"
)

type Character struct {
	nom    string
	classe string
	niveau int
	pv_max int
	pv     int
	inventaire Inventory
	spell spellbook
}

func InitCharacter() Character {
	var player1 Character
	player1.nom = "Rudolf"
	player1.classe = "Elfe"
	player1.niveau = 1
	player1.pv_max = 100
	player1.pv = 40
	player1.inventaire = initInventory()
	player1.spell = initSpellbook()
	return player1
}

func IsDead(player1 Character) bool {
	var i int
	if player1.pv == 0 {
		Printfct("Vous etes mort dommage !!! ", 30, 0)
		fmt.Scanln(&i)
		player1 = InitCharacter()
		return true
	}
	return false
}

func DisplayInfo(player1 Character) {
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n Nom : %v\n Classe : %v\n Niveau : %v\n Pv_max : %v\n Pv : %v\n\n", player1.nom, player1.classe, player1.niveau, player1.pv_max, player1.pv)
}
