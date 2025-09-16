package red

import (
	"fmt"
	"time"
)

type spellbook struct {
	poing int
	dmg_poing int
}

type opponent struct {
	name string
	pv int
	dmg int
	reward int
}

func initSpellbook() spellbook {
	var spell spellbook
	spell.poing = 1
	spell.dmg_poing = 10
	return spell
}

func initMob() opponent {
	var mob opponent
	mob.name = "The Goblin"
	mob.dmg = 5
	mob.pv = 30
	mob.reward = 10
	return mob
}

func winfight(player Character, mob opponent) Character {
	var i int
	player.inventaire.piece_or += mob.reward
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nWell done, You defeat %v !\nYou gain %v gold", mob.name, mob.reward)
	Printfct("Wanna continue ?", 2, 3)
	fmt.Scanln(&i)
	return player
}

func opponentTurn(player1 Character, mob opponent) (Character, opponent) {
	player1.pv -= mob.dmg
	fmt.Printf("\nAH ! %v attack and you take %v damages\n", mob.name, mob.dmg)
	time.Sleep(1 *time.Second)
	fmt.Printf("\nYou Have %v hp left, %v have %v left\n", player1.pv, mob.name, mob.pv)
	return player1, mob
}

func attack(player1 Character, mob opponent) (Character, opponent) {
	mob.pv -= 10
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nYou attack %v for %v damages", mob.name, 10)
	return player1, mob
}

func fight(player1 Character) Character {
	var i int
	mob := initMob()

	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n%v appears !\n", mob.name)
	Printfct("1: Attack\n2: Talk\n3: Use Potion\n0: Run Away", 1, 5)
	for {
		if IsDead(player1) {
			return player1
		}
		if mob.pv <= 0 {
			return winfight(player1, mob)
		}
		fmt.Scanln(&i)
		switch i {
		case 0:
			Printfct("You run away", 30, 0)
			time.Sleep(1 *time.Second)
			return player1
		case 1:
			player1, mob = attack(player1, mob)
		case 2:
			fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nIt seems that the enemy does not want to talk.")
		case 3:
			player1 = TakePot(player1)
		default:
			println("invalid")
		}
		time.Sleep(1 *time.Second)
		player1, mob = opponentTurn(player1, mob)
		time.Sleep(1 *time.Second)
		Printfct("1: Attack\n2: Talk\n3: Use Item\n0: Run Away", 1, 8)
	}
}