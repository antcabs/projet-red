package main

import (
    "fmt"
)

func actionA() {
    fmt.Println("La touche 'a' a été pressée.")
    
    grass()
}

func actionB() {
    fmt.Println("La touche 'b' a été pressée.")
    // Insérez ici le code à exécuter lorsque la touche 'b' est pressée.
	mer()
}

func actionC() {
    fmt.Println("La touche 'c' a été pressée.")
    // Insérez ici le code à exécuter lorsque la touche 'c' est pressée.
    jeu()
}
func actionD() {
    fmt.Println("la touche 'd' a été pressée")
    Invent()
}
func actionE(){
    fmt.Println("la touche e a été pressée")
    creation()
}
func actionF() {
    fmt.Println("la touche f a été pressé")
    forge()
}
func actionG() {
	fmt.Println("la touche g a été pressé")
	main()
}
func actionH() {
	fmt.Println("la touche h a été pressé")
	return}
func actionI() {
    fmt.Println("bienvenu sur notre cours en terre battue")
    clay()
}
func actionJ() {
    fmt.Println("bienvenu sur notre cours magic")
    magic()
}
func game() {

	

    fmt.Println( "a permet de jouer ")
    fmt.Println( "b pour acceder a Ynovshop ")
    fmt.Println( "c pour acceder a la carte ")
    fmt.Println( "d pour acceder a son inventaire ")
    fmt.Println( "e pour créer son personnage ")
    fmt.Println( "f pour acceder a la forge ")
    fmt.Println( "g pour retourner en arrière ")
	fmt.Println("h pour retourner au menu")
    fmt.Println("i pour le level 2")
    fmt.Println("j pour le level 3")
    

    var keyPressed string
    _, err := fmt.Scan(&keyPressed)

    if err != nil {
        fmt.Println("Erreur lors de la lecture de la touche:", err)
        return
    }

    switch keyPressed {
    case "a":
        actionA()
    case "b":
        actionB()
    case "c":
        actionC()
    case "d":
        actionD()
    case "e":
        actionE()
    case "f":
        actionF()
	case "g":
		actionG()
	case "h":
		actionH()
    case "i":
        actionI()
    case "j":
        actionJ()
    
    default:
        fmt.Println("Touche non reconnue.")
    }
}