package red

import (
	"fmt"
	"math/rand"
	"time"
)

type spellbook struct {
	poing        int
	boule_de_feu int
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
	dmg         int
	dmg_emp     int
	dmg_spe     int
	dmg_spe_emp int
	taken_emp   int
	luck        int
}

func initStat() stat {
	var st stat
	st.dmg = 10
	st.dmg_emp = 100
	st.dmg_spe = 10
	st.dmg_spe_emp = 100
	st.taken_emp = 100
	st.luck = 100
	return st
}

func initSpellbook() spellbook {
	var spell spellbook
	spell.poing = 1
	spell.boule_de_feu = 0
	return spell
}

func initMob() opponent {
	ran := rand.Intn(100)
	var mob opponent
	if ran < 50 {
		mob.name = "Le Gobelin"
		mob.dmg = 5
		mob.pv = 30
		mob.reward = 10
		mob.exp = 10
	}
	if ran >= 50 && ran < 80 {
		mob.name = "L'Orque"
		mob.dmg = 10
		mob.pv = 40
		mob.reward = 20
		mob.exp = 20
	}
	if ran >= 80 && ran < 99 {
		mob.name = "Krenko, le roi des gobelins"
		mob.dmg = 20
		mob.pv = 100
		mob.reward = 100
		mob.exp = 50
	}
	if ran == 99 {
		mob.name = "Ureni, le dragon divin"
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
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nBravo, vous l'avez vaincu %v !\nVous gagnez %v or", mob.name, tune)
	player.exp += mob.exp
	if player.exp >= 100 {
		player = lvlup(player)
	}
	Printfct("Vous voulez continuer ?", 2, 7)
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
		fmt.Printf("\nAH ! %v attaque et tu prends %v dégâts. C'est un coup critique !\n", mob.name, dmg)
	} else {
		fmt.Printf("\n%v attaque et tu prends %v dégâts.\n", mob.name, dmg)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("\nVous avez %v pv restant, %v as %v pv restant.\n", player1.pv, mob.name, mob.pv)
	return player1, mob
}

func attack(player1 Character, mob opponent) (Character, opponent) {
	var i int

	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n1: Coup de poing\n")
	if (player1.spell.boule_de_feu == 1) {
		fmt.Printf("2: Boule de feu\n")
	}
	Printfct("\n", 0, 5 + player1.spell.boule_de_feu)
	for {
		fmt.Scanln(&i)
		switch i {
		case 1:
			return poing(player1, mob)
		case 2:
			if (player1.spell.boule_de_feu == 1) {
				return bdf(player1, mob)
			}
		default:
			println("invalid")
		}
	}
}

func bdf(player1 Character, mob opponent) (Character, opponent) {
	if player1.mana < 5 {
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nTu n'as pas assez de mana")
		return player1, mob
	}
	player1.mana -= 5
	dmg := (player1.st.dmg_spe * player1.st.dmg_spe_emp) / 100 + 2
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
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nTu attaques %v pour %v dégats", mob.name, dmg)
	} else {
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nAH ! Tu attaques %v pour %v dégats. C'est un coup critique !", mob.name, dmg)
	}
	return player1, mob
}

func poing(player1 Character, mob opponent) (Character, opponent) {
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
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nTu attaques %v pour %v dégats", mob.name, dmg)
	} else {
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nAH ! Tu attaques %v pour %v dégats. C'est un coup critique !", mob.name, dmg)
	}
	return player1, mob
}

func fight(player1 Character) Character {
	var i int
	mob := initMob()

	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n%v apparaît !\n", mob.name)
	Printfct("1: Attaquer\n2: Parler\n3: Utiliser une potion\n0: Fuir", 1, 3)
	for {
		if IsDead(player1) {
			return *InitCharacter()
		}
		fmt.Scanln(&i)
		switch i {
		case 0:
			Printfct("Tu t'enfuis", 30, 0)
			time.Sleep(1 * time.Second)
			return player1
		case 1:
			player1, mob = attack(player1, mob)
		case 2:
			Printfct("Il semble que l'ennemi ne veuille pas discuter.", 30, 0)
		case 3:
			player1 = TakePot(player1)
		case 12:
			Printfct("CHEAT CODE POWER !", 30, 0)
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
		Printfct("1: Attaquer\n2: Parler\n3: Utiliser un objet\n0: Fuir", 1, 8)
	}
}
