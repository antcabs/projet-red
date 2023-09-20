package main

import (
	"bufio"
	"fmt"
	"os"

)

type Personnage struct {
	nom string
	class string
	level int
	Hp int

}

func creation() Personnage{
	var perso Personnage

	fmt.Println("Création d un personnage:")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nom :")
	nom, _ := reader.ReadString('\n')
	perso.nom = nom

	fmt.Print("Class(serveur,défenseur,attaquant): ")
	class, _ := reader.ReadString('\n')
	perso.class = class

	fmt.Print("Niveau :")
	fmt.Scan(&perso.level)

	switch perso.class{
	case "serveur":
		perso.Hp = 100
	case "défenseur":
		perso.Hp = 120
	case "attaquant":
		perso.Hp = 80
	default:
		fmt.Println("Class invalide. Point de vie par défaut: 100")
		perso.Hp = 100
	}

	return perso
}
