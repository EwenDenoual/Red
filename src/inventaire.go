package red

import "fmt"

type Inventory struct {
	possion int
}

func initInventory() Inventory {
	var inventaire_player1 Inventory
	inventaire_player1.possion = 3
	return inventaire_player1
}


func AccessInventory() {
	inventaire1_player1 := initInventory()
	fmt.Printf("\n Possion = %v\n", inventaire1_player1.possion)
}
