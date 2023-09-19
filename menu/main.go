package main

import "fmt"

type Character struct {
	Name string
	Height int
	Weight int
	Haircolor string
	Level int
	Hp int
	Power int
	Speed int
	Money int
}

type Inventory struct {
	Items []string
}

func main() {
	character := Character{
		Name: "Gasan",
		Height: 180,
		Weight: 70,
		Haircolor: "Brun",
		Level: 1,
		Hp: 50,
		Power: 1,
		Speed: 10,
		Money: 0,
	}

	inventory := Inventory{
		Items: []string{"racket", "balls", "water"},
	}
	for{
		fmt.Println("Menu:")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à lnventaire")
		fmt.Println("3. Quitter")

		var option int
		fmt.Println("Choisissez une option: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("informations du personnage :")
			fmt.Println("Nom :", character.Name)
			fmt.Println("Taille: ", character.Height)
			fmt.Println("Poids: ", character.Weight)
			fmt.Println("couleur de cheveux:", character.Haircolor)
			fmt.Println("Niveau :", character.Level)
			fmt.Println("Points de vie :", character.Hp)
			fmt.Println("Attaque: ", character.Power)
			fmt.Println("Defense :", character.Speed)
			fmt.Println("Argent :", character.Money)
		case 2:
			fmt.Println("Inventaire :")
			for _, item := range inventory.Items{
				fmt.Println("-", item)
			}
		case 3:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("option invalide")
		}

		fmt.Println()
	}
}
