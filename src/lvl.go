package red

import (
	"fmt"
	"math/rand"
)

func lvlup(player Character) Character {
	roll := 0
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nWhouah tu as gagné un niveau !!!")
	player.exp -= 100
	player.niveau++
	roll = rand.Intn(11)
	player.pv_max += roll
	player.pv += roll
	fmt.Printf("\nPv max: +%v => %v\n", roll, player.pv_max)
	roll = rand.Intn(5)
	player.st.dmg += roll
	fmt.Printf("att: +%v => %v\n", roll, player.st.dmg)
	if player.classe == "Nains" {
		roll = rand.Intn(3)
	} else if player.classe == "Elfes" {
		roll = rand.Intn(8)
	} else {
		roll = rand.Intn(5)
	}
	player.st.dmg_spe += roll
	fmt.Printf("att spe: +%v => %v\n", roll, player.st.dmg_spe)
	roll = rand.Intn(10) - 3
	player.st.luck += roll
	fmt.Printf("cha: +%v => %v\n", roll, player.st.luck)
	return player
}
