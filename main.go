package main

import (
	"red/src"
	"fmt"
	"time"
)

func config() {
	fmt.Print("\n\n\n\n\n\n")
	fmt.Printf("------------------------------------------------------------------------------------------------")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("POUR CONFIGURER VOTRE TERMINAL : AJUSTEZ SA TAILLE POUR QU’ELLE SOIT JUSTE AU-DESSUS DE LA BARRE")
	fmt.Print("\n--------------------------- UNE FOIS CONFIGURE CLIQUER 0 --------------------------------------")
	fmt.Print("\n\n\n\n\n\n\n\n\n")
	var entrer int
	fmt.Scanln(&entrer)
	if entrer == 0 {
		chargement()
	}
}

func chargement() {
	gobelin := `
       ,      ,
      /(.-""-.)\ 
  |\  \/      \/  /|
  | \ / =.  .= \ / |
  \( \   o\/o   / )/
   \_, '-/  \-' ,_/
     /   \__/   \
     \ \__/\__/ /
   ___\ \|--|/ /___
 /` + "`" + `    \      /    ` + "`" + `\
/       '----'       \
`
    fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Print(gobelin)
	fmt.Print("\n\n\n")
	fmt.Print("\n\n")
	charge := "="
	fmt.Printf("-- JEU EN COURS DE CHARGEMENT : --\n\n")
	for i := 0; i < 7; i++ {
		charge += "="
		fmt.Print(charge)
		time.Sleep(250 *time.Millisecond)
	}
	fmt.Print("\n\n")
	fmt.Print("----- VOUS ENTREZ DANS LE JEU -----\n")
	fmt.Print("\n\n")
	time.Sleep(2 *time.Second)
}

func main() {
	config()
	Character := red.InitCharacter()

	red.Choice(Character)
	for range 20 {
		fmt.Printf("\n")
	}
}
