package main

import "fmt"

func AccessInventory(inventory []string) {
    fmt.Println("Voici les items présents dans votre inventaire :")
    for _, item := range inventory {
        fmt.Println(item)
    }
}
    //menu permettant de selectionner 3 choix renvoyant vers des fonctions
func ReturnToMenu() {
    fmt.Println("Retour au menu principale")
}
func Invent() {
    inventory := []string{"Racket", "Balls", "Water"}

    fmt.Println("Menu :")
    fmt.Println("1. Acceder au jeu")
    fmt.Println("2. Afficher les donées du personnage")
    fmt.Println("3. Quitter")

    var choice int
    fmt.Scanln(&choice)
    //on appuie sur 1, 2 ou 3 pour seletionner si on va jouer, aller dans l inventaire ou quitter le jeu
    switch choice {
    case 1:
        fmt.Println("Bienvenue à Ynov Garros")
        main()
    case 2:
        AccessInventory(inventory)
    case 3:
        fmt.Println("Vous quittez le jeu.")
        // Autres actions pour quitter le jeu
    default:
        fmt.Println("Choix invalide.")
    }
}
