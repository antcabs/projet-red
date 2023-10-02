package main

import "fmt"

const maxInventoryCapacity = 10



func addItemToInventory(player *Player, item string) string {
    if len(player.inventory) < maxInventoryCapacity {
        player.inventory = append(player.inventory, item)
        return fmt.Sprintf("L'item '%s' a été ajouté à ton sac.", item)
    }
    return "Ton sac est déjà plein ! Impossible d'ajouter plus d'items."
}

func max() {
    player := Player{}

    // Ajoute 10 items à l'inventaire
    for i := 1; i <= maxInventoryCapacity; i++ {
        message := addItemToInventory(&player, fmt.Sprintf("Item %d", i))
        fmt.Println(message)
    }

    // pour essayer d'en ajouter un autre
    message := addItemToInventory(&player, "Item supplémentaire")
    fmt.Println(message)
}
