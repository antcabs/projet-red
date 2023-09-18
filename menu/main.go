package main

import "fmt"

type Character struct {
	Name    string
	Level   int
	HP      int
	Attack  int
	Defense int
}

type Inventory struct {
	Items []string
}

func main() {
	character := Character{
		Name:    "John",
		Level:   1,
		HP:      100,
		Attack:  5,
		Defense: 0,
	}

	inventory := Inventory{
		Items: []string{"racket", "Shield", "water"},
	}

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Quitter")

		var option int
		fmt.Print("Choisissez une option : ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Informations du personnage :")
			fmt.Println("Nom :", character.Name)
			fmt.Println("Niveau :", character.Level)
			fmt.Println("Points de vie :", character.HP)
			fmt.Println("Attaque :", character.Attack)
			fmt.Println("Défense :", character.Defense)
		case 2:
			fmt.Println("Inventaire :")
			for _, item := range inventory.Items {
				fmt.Println("-", item)
			}
		case 3:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Option invalide")
		}

		fmt.Println()
	}
}
