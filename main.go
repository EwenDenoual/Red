package main

import (
	"fmt"
	"red/src"
)


func main() {
	Character := red.InitCharacter()
	red.DisplayInfo(Character)
	var info int
	fmt.Print("Voulez vous utiliser une potion si oui, entrer le nombre de potion que vous voulez utiliser ")
	fmt.Scanln(&info)

	if info == 1 {
		red.TakePot(Character)
	}
}
