package main

import (
    "fmt"
)

func actionP() {
    fmt.Println("La touche 'p' à été pressée.")
    
    grass()
}
// switch case permettant d accéder à nos différentes fonctions
func actionS() {
    fmt.Println("La touche 's' à été pressée.")
    
	mer()
}

func actionC() {
    fmt.Println("La touche 'c' à été pressée.")
    
    jeu()
}
func actionD() {
    fmt.Println("la touche 'd' à été pressée")
    Invent()
}
func actionE(){
    fmt.Println("la touche e à été pressée")
    creation()
}
func actionF() {
    fmt.Println("la touche f à été pressé")
    forge()
}
func actionG() {
	fmt.Println("la touche g à été pressé")
	main()
}
func actionH() {
	fmt.Println("la touche h à été pressé")
	return}
func actionB() {
    fmt.Println("bienvenu sur notre cours en terre battue")
    clay()
}
func actionW() {
    fmt.Println("bienvenu sur notre cours magic")
    magic()
}
func game() {

    fmt.Println( "c pour acceder à la carte ")
    fmt.Println( "p pour acceder au premier niveau" )
    fmt.Println( "s pour acceder a Ynovshop")
    fmt.Println( "f pour acceder à la forge")
    fmt.Println( "b pour acceder au second niveau")
    fmt.Println( "w pour acceder au troisieme niveau")

    fmt.Println( "d pour acceder à son inventaire ")

    fmt.Println( "e pour créer son personnage ")

    fmt.Println( "g pour retourner en arrière ")

	fmt.Println( "h pour retourner au menu")
  
   
    

    var keyPressed string
    _, err := fmt.Scan(&keyPressed)

    if err != nil {
        fmt.Println("Erreur lors de la lecture de la touche:", err)
        return
    }

    switch keyPressed {
    case "p":
        actionP()
    case "s":
        actionS()
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
    case "b":
        actionB()
    case "w":
        actionW()
    
    default:
        fmt.Println("Touche non reconnue.")
    }
}