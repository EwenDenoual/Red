package main

import "fmt"

type Character struct {
	nom string 
	classe string 
	niveau int 
	pv_max int 
	pv int 
}

func initCharacter() Character {
	var player1 Character
	player1.nom = "Rudolf"
	player1.classe = "Elfe"
	player1.niveau = 1
	player1.pv_max = 100
	player1.pv = 40
	return player1
}

func main() {
	player1 := initCharacter()

	fmt.Printf("name : %v, classe : %v, niveau: %v\n", player1.nom, player1.classe, player1.niveau)
}
