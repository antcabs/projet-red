package main

import (
	"fmt"
	
	"time"
)

func wasted() {
	Character := Character{
		Name:      "Gasan",
		Health:    0,
		MaxHealth: 100,
		Damage:    10,
	}

	fmt.Printf("%s is dead. Resurrecting...\n", Character.Name)
	dead(&Character)

	fmt.Printf("%s's health: %d\n", Character.Name, Character.Health)
}

func dead(Character *Character) {
	if Character.Health <= 0 && !Character.Resurrected {
		Character.IsDead = true
		Character.Resurrected = true

		// Ressusciter avec 50% de points de vie max
		Character.Health = Character.MaxHealth / 2

		// Afficher un message de résurrection
		fmt.Printf("%s has been resurrected with %d health!\n", Character.Name, Character.Health)

		// Ajouter un délai pour l'effet dramatique :)
		time.Sleep(1 * time.Second)
	}
}