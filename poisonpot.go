package main

import (
    "fmt"
    "time"
)

func poisonPot(maxFatigue int) {
    currentHP := maxFatigue
    fmt.Printf("Fatigue actuelle: %d/%d\n", currentHP, maxFatigue)

    for i := 1; i <= 3; i++ { // Les dégâts pendant 3 secondes
        currentHP -= 10
        fmt.Printf("Usure de la semelle pendant la seconde %d : -10 PV\n", i)
        fmt.Printf("Points de fatigue actuelle: %d/%d\n", currentHP, maxFatigue)

        time.Sleep(time.Second) // Pour le délai d'une seconde
    }
}

func poison() {
    maxFatigue := 100 // Points de vie maximum
    poisonPot(maxFatigue)
}
