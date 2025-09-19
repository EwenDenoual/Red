package red

import (
	"os"
	"strconv"
)

func get_line_string(data string, n int) string {
	cmp_n := 0
	ret := ""
	for _, v := range data {
		if v == '\n' {
			cmp_n++
		} else if cmp_n == n {
			ret += string(v)
		}
	}
	return ret
}

func get_line_int(data string, n int) int {
	cmp_n := 0
	ret := 0
	for _, v := range data {
		if v == '\n' {
			cmp_n++
		} else if cmp_n == n {
			ret = ret * 10 + int(v - '0')
		}
	}
	return ret
}

func Getsave() (*Character) {
	data, err := os.ReadFile(".save")
	if err != nil {
		file, _ :=os.OpenFile(".save", os.O_CREATE, 0600)
		file.Close()
		return nil
	}
	var chara Character
	chara.nom = get_line_string(string(data), 0)	
	chara.classe = get_line_string(string(data), 1)
	chara.niveau = get_line_int(string(data), 2)
	chara.exp = get_line_int(string(data), 3)
	chara.pv_max = get_line_int(string(data), 4)
	chara.pv = get_line_int(string(data), 5)
	chara.mana = get_line_int(string(data), 6)

	chara.st.dmg = get_line_int(string(data), 7)
	chara.st.dmg_emp = get_line_int(string(data), 8)
	chara.st.taken_emp = get_line_int(string(data), 9)
	chara.st.luck = get_line_int(string(data), 10)

	chara.inventaire.potion = get_line_int(string(data), 11)
	chara.inventaire.potion_poison = get_line_int(string(data), 12)
	chara.inventaire.piece_or = get_line_int(string(data), 13)

	chara.inventaire.Plume_de_Corbeau = get_line_int(string(data), 14)
	chara.inventaire.Cuir_de_Sanglier = get_line_int(string(data), 15)
	chara.inventaire.Fourrure_de_loup = get_line_int(string(data), 16)
	chara.inventaire.Peau_de_Troll = get_line_int(string(data), 17)

	chara.inventaire.Chapeau_de_l_aventurier = get_line_int(string(data), 18)
	chara.inventaire.Tunique_de_l_aventurier = get_line_int(string(data), 19)
	chara.inventaire.Bottes_de_l_aventurier = get_line_int(string(data), 20)

	chara.inventaire.size_max = get_line_int(string(data), 21)

	chara.spell.poing = get_line_int(string(data), 22)
	chara.spell.boule_de_feu = get_line_int(string(data), 23)

	chara.equipment.equipement_tete = get_line_int(string(data), 24)
	chara.equipment.equipement_torse = get_line_int(string(data), 25)
	chara.equipment.equipement_pieds = get_line_int(string(data), 26)
	return &chara
}

func Savegame(chara Character) {
	file, _ :=os.OpenFile(".save", os.O_WRONLY, 0600)

	file.WriteString(chara.nom + "\n")
	file.WriteString(chara.classe + "\n")
	file.WriteString(strconv.Itoa(chara.niveau) + "\n")
	file.WriteString(strconv.Itoa(chara.exp) + "\n")
	file.WriteString(strconv.Itoa(chara.pv_max) + "\n")
	file.WriteString(strconv.Itoa(chara.pv) + "\n")
	file.WriteString(strconv.Itoa(chara.mana) + "\n")

	file.WriteString(strconv.Itoa(chara.st.dmg) + "\n")
	file.WriteString(strconv.Itoa(chara.st.dmg_emp) + "\n")
	file.WriteString(strconv.Itoa(chara.st.taken_emp) + "\n")
	file.WriteString(strconv.Itoa(chara.st.luck) + "\n")

	file.WriteString(strconv.Itoa(chara.inventaire.potion) + "\n")
	file.WriteString(strconv.Itoa(chara.inventaire.potion_poison) + "\n")
	file.WriteString(strconv.Itoa(chara.inventaire.piece_or) + "\n")

	file.WriteString(strconv.Itoa(chara.inventaire.Plume_de_Corbeau) + "\n")
	file.WriteString(strconv.Itoa(chara.inventaire.Cuir_de_Sanglier) + "\n")
	file.WriteString(strconv.Itoa(chara.inventaire.Fourrure_de_loup) + "\n")
	file.WriteString(strconv.Itoa(chara.inventaire.Peau_de_Troll) + "\n")

	file.WriteString(strconv.Itoa(chara.inventaire.Chapeau_de_l_aventurier) + "\n")
	file.WriteString(strconv.Itoa(chara.inventaire.Tunique_de_l_aventurier) + "\n")
	file.WriteString(strconv.Itoa(chara.inventaire.Bottes_de_l_aventurier) + "\n")

	file.WriteString(strconv.Itoa(chara.inventaire.size_max) + "\n")
	
	file.WriteString(strconv.Itoa(chara.spell.poing) + "\n")
	file.WriteString(strconv.Itoa(chara.spell.boule_de_feu) + "\n")
	
	file.WriteString(strconv.Itoa(chara.equipment.equipement_tete) + "\n")
	file.WriteString(strconv.Itoa(chara.equipment.equipement_torse) + "\n")
	file.WriteString(strconv.Itoa(chara.equipment.equipement_pieds) + "\n")
}

// ./save structure :

// name string
// classe string
// niveau int
// exp int
// pv_max int
// pv int

// st
// 	dmg int
// 	dmg_emp int
// 	taken_emp int
// 	luck int

// inventaire
// 	potion int
// 	potion_poison int
// 	piece_or int
	
// 	Plume_de_Corbeau int
// 	Cuir_de_Sanglier int
// 	Fourrure_de_loup int
// 	Peau_de_Troll int

// 	Chapeau_de_l_aventurier int
// 	Tunique_de_l_aventurier int
// 	Bottes_de_l_aventurier int

//  Size_max int

// spell
// 	poing     int
// 	boule_de_feu int

// equipment
// 	equipement_tete int
// 	equipement_torse int
// 	equipement_pieds int