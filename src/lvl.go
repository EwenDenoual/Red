package red

import (
	"fmt"
	"math/rand"
)

func lvlup(player Character) Character {
	roll := 0
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nWhouah you level up !!!")
	player.exp -= 100
	player.niveau++
	roll = rand.Intn(11)
	player.pv_max += roll
	player.pv += roll
	fmt.Printf("\npv max: +%v => %v\n", roll, player.pv_max)
	roll = rand.Intn(5)
	player.st.dmg += roll
	fmt.Printf("att: +%v => %v\n", roll, player.st.dmg)
	roll = rand.Intn(10) - 3
	player.st.luck += roll
	fmt.Printf("cha: +%v => %v\n", roll, player.st.luck)
	return player
}