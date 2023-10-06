package main

import "fmt"

func ptvie() {
	MaxHealth := 100

	casquetteEquipe := true
	TeeshirtEquipe := true
	TennisEquipe := true

	if casquetteEquipe{
		MaxHealth += 10
	}
	if TeeshirtEquipe{
		MaxHealth += 25
	}
	if TennisEquipe{
		MaxHealth += 15
	}
	fmt.Println("Le point de vie max du personnage est de", MaxHealth)
}
