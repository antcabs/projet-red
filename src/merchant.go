package main

import (
	"fmt"
)

type Item struct {
	Name string
	Price int
}

type Merchant struct {
	Items []Item
}


func mer() {
	menu := []string{"Bienvenue chez ynovshop !"}
	inventaire := []Item{
		{Name: "Pièces", Price: 10},
	}
	merchant := Merchant{
		Items: []Item{
			{Name: "Eau", Price: 0},
			{Name: "Wooden racket", Price: 10},
			{Name: "Iron racket", Price: 25},
			{Name:"Steel racket", Price: 50},
			{Name:"Carbon racket", Price: 75},
			{Name:"Titanium racket", Price: 100},
			{Name:"Diamont racket", Price: 200},
			{Name:"Red Powerade", Price: 10},
			{Name:"blue Powerade", Price: 10},
			
		},
	}

	for {
		fmt.Println("Menu:")
		for i, choix := range menu{
			fmt.Printf("%d.%s\n",i+1,choix)
		}
		var choixMenu int
		fmt.Print("Appuyer sur 1 pour acheter / 2 pour quitter")
		fmt.Scan(&choixMenu)
		switch choixMenu {
		case 1:
			fmt.Println("Bienvenue dans la boutique Ynovtennis!")
			
			for i, item := range
		merchant.Items{
			fmt.Printf("%d.%s(Price:%d coins)\n",i+1,item.Name,item.Price)
		}
		
		var choixItem int
		fmt.Print("choisissez un item à acheter")
		fmt.Scan(&choixItem)

		if choixItem > 0 && choixItem <= len(merchant.Items) {
			itemAchete := merchant.Items[choixItem-1]
			pieces := getNumberOfCoins(inventaire)
			if itemAchete.Price <= pieces {
			inventaire = append(inventaire, itemAchete)
			inventaire[0].Price -= itemAchete.Price
			fmt.Printf("Vous avez acheté: %s\n", itemAchete.Name)
			fmt.Printf("il vous reste %d pièces.\n", inventaire[0].Price)
			fmt.Println("2 pour revenir au menu")
		}else{
			fmt.Println("Vous nvez pas assez de pièces pour acheter cet item.")
		}
	}else{
		fmt.Println("choix invalide.")
	}
	case 2:
		fmt.Println("Au revoir !")
		return

	default:
		fmt.Println("choix invalide.")
		}
		fmt.Println("")
		}
	}
	func getNumberOfCoins(inventaire[]Item) int {
	return inventaire[0].Price
	}
