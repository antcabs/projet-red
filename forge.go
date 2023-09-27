package main

import (
    "fmt"
)

type Ressources struct {
    CarbonRacket   int
    TitaniumRacket int
}

type Weapon struct {
    name  string
    power int
}

type Forgeron struct {
    availableWeapons map[string]int
}

func NewForgeron() *Forgeron {
    return &Forgeron{
        availableWeapons: map[string]int{
            "Raquette hantée": 10,
            "Balles en or":    8,
            "Jetpack":         12,
        },
    }
}

func (f *Forgeron) ShowAvailableWeapons() {
    fmt.Println("Items spéciaux disponibles:")
    for weapon, power := range f.availableWeapons {
        fmt.Printf("%s - Puissance: %d\n", weapon, power)
    }
}

func (f *Forgeron) FabriquerArme(ressources Ressources, arme string) (Weapon, error) {
    if power, ok := f.availableWeapons[arme]; ok {
        if ressources.CarbonRacket >= 5 && ressources.TitaniumRacket >= 3 {
            // Fabrication réussie
            ressources.CarbonRacket -= 5
            ressources.TitaniumRacket -= 3

            return Weapon{name: arme, power: power}, nil
        } else {
            return Weapon{}, fmt.Errorf("Items insuffisants pour fabriquer %s", arme)
        }
    } else {
        return Weapon{}, fmt.Errorf("L'item %s n'est pas disponible chez ton vestiaire", arme)
    }
}

func forge() {
    forgeron := NewForgeron()
    ressources := Ressources{CarbonRacket: 10, TitaniumRacket: 7}

    fmt.Println("Te voilà dans ton vestiaire !")
    forgeron.ShowAvailableWeapons()

    // Exemple de fabrication des balles en or
    arme, err := forgeron.FabriquerArme(ressources, "Balles en or")
    if err != nil {
        fmt.Println("Erreur:", err)
    } else {
        fmt.Printf("Item fabriqué: %s (Puissance: %d)\n", arme.name, arme.power)
        fmt.Printf("Items restantes: Carbon Racket - %d, Titanium Racket - %d\n", ressources.CarbonRacket, ressources.TitaniumRacket)
    }
}
