package main

import (
	"red/src"
	"fmt"
)

func main() {
	Character := red.InitCharacter()
	red.Choice(Character)
	for range 20 {
		fmt.Printf("\n")
	}
	// red.DisplayInfo(Character)
	// var info int
	// fmt.Print("Voulez vous utiliser une potion si oui, entrer le nombre de potion que vous voulez utiliser ")
	// fmt.Scanln(&info)

	// if info == 1 {
	// 	red.TakePot(Character)
	// }

	// if info == 2 {
	// 	red.PoisonPot(Character)
	// }
	
}
