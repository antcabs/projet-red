package main

import "fmt"

var upgradeCount = 3
var maxCapacity = 100

func upgradeInventorySlot() string {
    if upgradeCount > 0 {
        maxCapacity += 10
        upgradeCount--
        return fmt.Sprintf("T'es bien drip, quel style ! Ta jauge de fatigue est désormais à %d !", maxCapacity)
    } else {
        return fmt.Sprintf("Nombre maximal d'améliorations atteint. Capacité maximale inchangée : %d", maxCapacity)
    }
}
