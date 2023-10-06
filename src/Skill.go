package main

import "fmt"


func assignSkills(character*Character, spells[] string) {
	for _, spell := range spells {
		character.Skills = append(character.Skills, spell)
	}
}

func spellbook(character*Character, spellToAdd string) {
	for _, Skills := range character.Skills {
		if Skills == spellToAdd{
			return
		}
	}
	character.Skills = append(character.Skills, spellToAdd)
}

func skill() {
	spells := []string{"coup droit", ""}
	player := Character{Name: "Gasan", Skills:[]string{}}

	assignSkills(&player, spells)
	
	fmt.Println(player)
}
