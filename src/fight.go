package red

import (
	"fmt"
	"math/rand"
	"time"
)

type spellbook struct {
	poing     int
	dmg_poing int
}

type opponent struct {
	name   string
	pv     int
	dmg    int
	reward int
	exp    int
	crit   int
}

type stat struct {
	dmg       int
	dmg_emp   int
	taken_emp int
	luck      int
}

func initStat() stat {
	var st stat
	st.dmg = 10
	st.dmg_emp = 100
	st.taken_emp = 100
	st.luck = 100
	return st
}

func initSpellbook() spellbook {
	var spell spellbook
	spell.poing = 1
	spell.dmg_poing = 10
	return spell
}

func initMob() opponent {
	ran := rand.Intn(100)
	var mob opponent
	if ran < 50 {
		mob.name = "The Goblin"
		mob.dmg = 5
		mob.pv = 30
		mob.reward = 10
		mob.exp = 10
	}
	if ran >= 50 && ran < 80 {
		mob.name = "The Orque"
		mob.dmg = 10
		mob.pv = 40
		mob.reward = 20
		mob.exp = 20
	}
	if ran >= 80 && ran < 99 {
		mob.name = "Krenko, The Goblin King"
		mob.dmg = 20
		mob.pv = 100
		mob.reward = 100
		mob.exp = 50
	}
	if ran == 99 {
		mob.name = "Ureni, The Divin Dragon"
		mob.dmg = 99
		mob.pv = 10000
		mob.reward = 99999999999
		mob.exp = 100
	}
	mob.crit = 0
	return mob
}

func winfight(player Character, mob opponent) Character {
	var i int
	tune := (mob.reward * player.st.luck) / 100
	player.inventaire.piece_or += tune
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nWell done, You defeat %v !\nYou gain %v gold", mob.name, tune)
	player.exp += mob.exp
	if player.exp >= 100 {
		player = lvlup(player)
	}
	Printfct("Wanna continue ?", 2, 6)
	fmt.Scanln(&i)
	return player
}

func opponentTurn(player1 Character, mob opponent) (Character, opponent) {
	dmg := (mob.dmg * player1.st.taken_emp) / 100
	mob.crit++
	if mob.crit == 3 {
		dmg *= 2
		mob.crit = 0
	}
	player1.pv -= dmg
	if player1.pv < 0 {
		player1.pv = 0
	}
	if mob.crit == 0 {
		fmt.Printf("\nAH ! %v attack and you take %v damages. It's a critical hit !\n", mob.name, dmg)
	} else {
		fmt.Printf("\n%v attack and you take %v damages\n", mob.name, dmg)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("\nYou Have %v hp left, %v have %v left\n", player1.pv, mob.name, mob.pv)
	return player1, mob
}

func attack(player1 Character, mob opponent) (Character, opponent) {
	dmg := (player1.st.dmg * player1.st.dmg_emp) / 100
	crit := rand.Intn(100) + 100
	if crit < player1.st.luck {
		crit = 1
		dmg *= 2
	}
	mob.pv -= dmg
	if mob.pv < 0 {
		mob.pv = 0
	}
	if crit != 1 {
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nYou attack %v for %v damages", mob.name, dmg)
	} else {
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nAH ! You attack %v for %v damages. It's a critical hit !", mob.name, dmg)
	}
	return player1, mob
}

func fight(player1 Character) Character {
	var i int
	mob := initMob()

	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n%v appears !\n", mob.name)
	Printfct("1: Attack\n2: Talk\n3: Use Potion\n0: Run Away", 1, 3)
	for {
		if IsDead(player1) {
			return InitCharacter()
		}
		fmt.Scanln(&i)
		switch i {
		case 0:
			Printfct("You run away", 30, 0)
			time.Sleep(1 * time.Second)
			return player1
		case 1:
			player1, mob = attack(player1, mob)
		case 2:
			fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nIt seems that the enemy does not want to talk.")
		case 3:
			player1 = TakePot(player1)
		case 12:
			fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nCHEAT CODE POWER !")
			mob.pv = 0
		default:
			println("invalid")
		}
		time.Sleep(1 * time.Second)
		if mob.pv <= 0 {
			return winfight(player1, mob)
		}
		player1, mob = opponentTurn(player1, mob)
		time.Sleep(1 * time.Second)
		Printfct("1: Attack\n2: Talk\n3: Use Item\n0: Run Away", 1, 8)
	}
}
