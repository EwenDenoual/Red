package red

func lvlup(player Character) Character {
	player.exp -= 100
	player.niveau++
	return player
}