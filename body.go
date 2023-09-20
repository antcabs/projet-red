package main

import "fmt"



func equip() {
	equipement := Equipement{
		Head: "Casquette",
		Torso: "Tee shirt",
		Feet: "Tennis",
	}
	fmt.Printf("Equipement de tête: %s\n",equipement.Head)
	fmt.Printf("Equipement pour le : %s\n",equipement.Torso)
	fmt.Printf("eEquipement pour les pieds: %s\n", equipement.Feet)
}
